package sourcetest

import (
	"errors"

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
		// sort.Strings(scenario.Verif)
		// keys := src.keys()
		// sort.Strings(keys)
		// for _, k1 := range scenario.Verif {
		// 	for _, k2 := range src.keys() {

		// 		// col, err := src.Collection(karigo.QueryCol{
		// 		// 	Set: set,
		// 		// 	// Fields: []string{}, TODO Fields?
		// 		// 	// Sort: []string{}, TODO Sorting?
		// 		// 	PageSize: int(math.MaxInt64),
		// 		// })
		// 		// if err != nil {
		// 		// 	return err
		// 		// }
		// 		// if col == nil {
		// 		// 	return errors.New("no collection returned")
		// 		// }
		// 	}
		// }
	}

	return nil
}
