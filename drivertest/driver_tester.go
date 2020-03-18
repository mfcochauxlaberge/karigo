package drivertest

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/drivertest/internal/scenarios"
	"github.com/mfcochauxlaberge/karigo/internal/gold"
	"github.com/mfcochauxlaberge/karigo/query"

	"github.com/stretchr/testify/assert"
)

// Test ...
func Test(t *testing.T, src karigo.Source, jrnl karigo.Journal) {
	assert := assert.New(t)

	scenarios := scenarios.Scenarios

	runner := gold.NewRunner("testdata/goldenfiles/scenarios")

	err := runner.Prepare()
	if err != nil {
		panic(err)
	}

	// Run scenarios
	for _, scenario := range scenarios {
		// Reset
		err := src.Reset(karigo.FirstSchema())
		assert.NoError(err)

		tx, _ := src.NewTx()

		// Run each step
		for _, step := range scenario.Steps {
			switch s := step.(type) {
			case karigo.Op:
				ss := []karigo.Op{s}

				err := tx.Apply(ss)
				assert.NoError(err)

				enc, err := karigo.Encode(0, ss)
				assert.NoError(err)

				err = jrnl.Append(enc)
				assert.NoError(err)

			case []karigo.Op:
				err := tx.Apply(s)
				assert.NoError(err)

				enc, err := karigo.Encode(0, s)
				assert.NoError(err)

				err = jrnl.Append(enc)
				assert.NoError(err)

			default:
				panic("karigo: unknown step")
			}
		}

		// Check
		// verif := scenario.Verif
		keys := []string{}

		// Keys
		sets, err := tx.Collection(query.Col{
			Set:        "0_sets",
			Sort:       []string{"id"},
			PageNumber: 0,
			PageSize:   1000,
		})
		assert.NoError(err)

		for i := 0; i < sets.Len(); i++ {
			set := sets.At(i)
			id := set.GetID()

			col, err := tx.Collection(query.Col{
				Set:        id,
				Sort:       []string{"id"},
				PageNumber: 0,
				PageSize:   1000,
			})
			assert.NoError(err)

			// For each resource...
			for j := 0; j < col.Len(); j++ {
				res := col.At(j)

				// Add a key for each attribute.
				for _, attr := range res.Attrs() {
					v := res.Get(attr.Name)
					key := fmt.Sprintf("%s.%s.%s", id, res.GetID(), attr.Name)
					keys = append(keys, fmt.Sprintf("%s=%v", key, v))
				}
				// Add a key for each relationsip.
				for _, rel := range res.Rels() {
					key := fmt.Sprintf("%s.%s.%s", id, res.GetID(), rel.FromName)

					var r string

					if rel.ToOne {
						r = res.GetToOne(rel.FromName)
					} else {
						rs := res.GetToMany(rel.FromName)
						r = strings.Join(rs, ",")
					}

					keys = append(keys, fmt.Sprintf("%s=%s", key, r))
				}
			}
		}

		// sort.Strings(verif)
		sort.Strings(keys)
		out := []byte(strings.Join(keys, "\n"))

		// Golden file
		filename := strings.Replace(scenario.Name, " ", "_", -1) + ".txt"

		err = runner.Test(filename, out)
		if _, ok := err.(gold.ComparisonError); ok {
			assert.Fail("file is different", scenario.Name)
		} else {
			assert.NoError(err)
		}

		// Test journal
		i, _, _ := jrnl.Newest()
		entries, _ := jrnl.Range(0, i)

		journalOut := []byte{}

		for _, entry := range entries {
			ops, err := karigo.Decode(0, entry)
			assert.NoError(err)

			for _, op := range ops {
				journalOut = append(journalOut, []byte(op.String())...)
				journalOut = append(journalOut, '\n')
			}
		}

		if len(journalOut) > 0 {
			journalOut = journalOut[:len(journalOut)-1]
		}

		filename = strings.Replace(scenario.Name, " ", "_", -1) + ".journal.txt"

		err = runner.Test(filename, journalOut)
		if _, ok := err.(gold.ComparisonError); ok {
			assert.Fail("file is different", scenario.Name)
		} else {
			assert.NoError(err)
		}
	}
}
