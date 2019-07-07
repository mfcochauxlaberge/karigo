package sourcetest

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/sourcetest/internal/scenarios"

	"github.com/stretchr/testify/assert"
)

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
		verif := scenario.Verif
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

		for _, set := range sets {
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
			for _, res := range col {
				// Add a key for each attribute.
				for _, attr := range res.Attrs() {
					v := res.Get(attr.Name)
					key := fmt.Sprintf("%s.%s.%s", id, res.GetID(), attr.Name)
					keys = append(keys, fmt.Sprintf("%s=%v", key, v))
				}
				// Add a key for each relationsip.
				for _, rel := range res.Rels() {
					key := fmt.Sprintf("%s.%s.%s", id, res.GetID(), rel.Name)
					var r string
					if rel.ToOne {
						r = res.GetToOne(rel.Name)
					} else {
						rs := res.GetToMany(rel.Name)
						r = strings.Join(rs, ",")
					}
					keys = append(keys, fmt.Sprintf("%s=%s", key, r))
				}
			}
		}

		sort.Strings(verif)
		sort.Strings(keys)

		assert.Equal(verif, keys)
	}

	return nil
}
