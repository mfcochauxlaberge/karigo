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
			PageSize:   0,
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
				PageSize:   0,
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

		errorFound := false
		i1 := 0
		i2 := 0
		for i1 < len(verif) || i2 < len(keys) {
			var key1, key2 string
			if i1 < len(verif) {
				key1 = verif[i1]
			}
			if i2 < len(keys) {
				key2 = keys[i2]
			}

			// Both lists have been fully read
			if key1 == "" && key2 == "" {
				break
			}

			// List of keys is too short
			if key1 != "" && key2 == "" {
				errorFound = true
				fmt.Printf("Key %q is missing.\n", key2)
				i1++
				continue
			}

			// List of keys is too long
			if key1 == "" && key2 != "" {
				errorFound = true
				fmt.Printf("Key %q is not supposed to exist.\n", key2)
				i2++
				continue
			}

			// Both keys are the same
			if key1 == key2 {
				i1++
				i2++
				continue
			}

			// Key is not in verification list
			if key1 < key2 {
				errorFound = true
				fmt.Printf("Key %q is missing.\n", key2)
				i1++
				continue
			}

			// Key from verification list does not exist
			if key1 > key2 {
				errorFound = true
				fmt.Printf("Key %q is not supposed to exist.\n", key2)
				i2++
				continue
			}
		}

		if errorFound {
			return errors.New("verification failed")
		}
	}

	return nil
}
