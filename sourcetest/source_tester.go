package sourcetest

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/sourcetest/internal/scenarios"
)

// Test ...
func Test(src karigo.Source) error {
	scenarios := scenarios.Scenarios

	// Run scenarios
	for _, scenario := range scenarios {
		fmt.Printf("A SCENARIO!\n")

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
			// case string:
			// 	switch s {
			// 	case "rollback":
			// 		err = srcRollback()
			// 		if err != nil {
			// 			return err
			// 		}
			// 	case "commit":
			// 		err = srcCommit()
			// 		if err != nil {
			// 			return err
			// 		}
			// 	}
			default:
				return errors.New("karigo: unknown step")
			}
		}

		// Check
		// TODO Check the result
		keys := []string{}

		sets, err := src.Collection(karigo.QueryCol{
			Set:        "0_sets",
			Sort:       []string{"id"},
			PageNumber: 0,
			PageSize:   0,
		})
		if err != nil {
			return fmt.Errorf("could not get list of sets: %s", err)
		}

		for i, set := range sets {
			id := set.GetID()

			fmt.Printf("Set %d: %s\n", i, id)

			col, err := src.Collection(karigo.QueryCol{
				Set:        id,
				Sort:       []string{"id"},
				PageNumber: 0,
				PageSize:   0,
			})
			if err != nil {
				return fmt.Errorf("could not get collection from %q: %s", id, err)
			}

			// For each resource...
			for _, res := range col {
				// fmt.Printf("Col!\n")
				// Add a key for each attribute.
				for _, attr := range res.Attrs() {
					// fmt.Printf("Attr!\n")
					v := res.Get(attr.Name)
					keys = append(keys, fmt.Sprintf("%s=%v", attr.Name, v))
				}
				// Add a key for each relationsip.
				for _, rel := range res.Rels() {
					if rel.ToOne {
						r := res.GetToOne(rel.Name)
						keys = append(keys, fmt.Sprintf("%s=%s", rel.Name, r))
					} else {
						rs := res.GetToMany(rel.Name)
						list := strings.Join(rs, ",")
						keys = append(keys, fmt.Sprintf("%s=%s", rel.Name, list))
					}
				}
			}
		}

		sort.Strings(scenario.Verif)
		sort.Strings(keys)

		fmt.Printf("Verif: %v\n", scenario.Verif)
		fmt.Printf("keys: %v\n", keys)

		i := 0
		j := 0
		for i = 0; i < len(scenario.Verif) && j < len(keys); {
			key1 := scenario.Verif[i]
			key2 := keys[j]

			if key1 == key2 {
				i++
				j++
				break
			}

			if key1 < key2 {
				fmt.Printf("Key %q is missing.", key1)
				i++
			} else {
				fmt.Printf("Key %q is not supposed to exist.", key1)
				j++
			}
		}
	}

	return nil
}
