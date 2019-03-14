package sourcetest

import (
	"github.com/mfcochauxlaberge/karigo"
)

// Test ...
func Test(src karigo.Source) error {
	// scenarios := scenarios.Scenarios

	// // Run scenarios
	// for _, scenario := range scenarios {
	// 	// Reset
	// 	err := src.Reset()
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Begin transaction
	// 	tx, err := src.Begin()
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Run each step
	// 	for _, step := range scenario.Steps {
	// 		switch s := step.(type) {
	// 		case karigo.Op:
	// 			err := tx.Apply([]karigo.Op{s})
	// 			if err != nil {
	// 				return err
	// 			}
	// 		case []karigo.Op:
	// 			err := tx.Apply(s)
	// 			if err != nil {
	// 				return err
	// 			}
	// 		case string:
	// 			switch s {
	// 			case "rollback":
	// 				err = tx.Rollback()
	// 				if err != nil {
	// 					return err
	// 				}
	// 			case "commit":
	// 				err = tx.Commit()
	// 				if err != nil {
	// 					return err
	// 				}
	// 			}
	// 		default:
	// 			return errors.New("karigo: unknown step")
	// 		}
	// 	}

	// 	// Check
	// 	for set := range scenario.Verif {
	// 		col, err := src.Collection(karigo.QueryCol{
	// 			Set: set,
	// 			// Fields: []string{}, TODO Fields?
	// 			// Sort: []string{}, TODO Sorting?
	// 			PageSize: int(math.MaxInt64),
	// 		})
	// 		if err != nil {
	// 			return err
	// 		}
	// 		if col == nil {
	// 			return errors.New("no collection returned")
	// 		}
	// 	}
	// }

	return nil
}
