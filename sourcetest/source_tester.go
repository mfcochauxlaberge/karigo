package sourcetest

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/sourcetest/internal/scenarios"

	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update-golden-files", false, "update the golden files")

// Test ...
func Test(t *testing.T, src karigo.Source) error {
	assert := assert.New(t)

	scenarios := scenarios.Scenarios

	// Run scenarios
	for _, scenario := range scenarios {
		// Reset
		err := src.Reset()
		if err != nil {
			return err
		}

		// Run each step
		for _, step := range scenario.Steps {
			switch s := step.(type) {
			case karigo.Op:
				err := src.Apply([]karigo.Op{s})
				if err != nil {
					return err
				}
			case []karigo.Op:
				err := src.Apply(s)
				if err != nil {
					return err
				}
			default:
				return errors.New("karigo: unknown step")
			}
		}

		// Check
		// verif := scenario.Verif
		keys := []string{}

		// Keys
		sets, err := src.Collection(karigo.QueryCol{
			Set:        "0_sets",
			Sort:       []string{"id"},
			PageNumber: 0,
			PageSize:   1000,
		})
		if err != nil {
			return fmt.Errorf("could not get list of sets: %s", err)
		}

		for i := 0; i < sets.Len(); i++ {
			set := sets.At(i)
			id := set.GetID()

			col, err := src.Collection(karigo.QueryCol{
				Set:        id,
				Sort:       []string{"id"},
				PageNumber: 0,
				PageSize:   1000,
			})
			if err != nil {
				return fmt.Errorf("could not get collection from %q: %s", id, err)
			}

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

		// Golden file
		filename := strings.Replace(scenario.Name, " ", "_", -1) + ".txt"
		path := filepath.Join("testdata", "goldenfiles", "scenarios", filename)

		if !*update {
			// Retrieve the expected result from file
			contents, _ := ioutil.ReadFile(path)
			expected := []string{}

			for _, key := range strings.Split(string(contents), "\n") {
				if key != "" {
					expected = append(expected, key)
				}
			}

			assert.Equal(expected, keys, scenario.Name)
		} else {
			dst := &bytes.Buffer{}
			for _, key := range keys {
				_, _ = fmt.Fprintln(dst, key)
			}

			// TODO Figure out whether 0644 is okay or not.
			err = ioutil.WriteFile(path, dst.Bytes(), 0644)
			assert.NoError(err)
		}
	}

	return nil
}
