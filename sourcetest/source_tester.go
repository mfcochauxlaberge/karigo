package sourcetest

import (
	"errors"

	"github.com/kkaribu/jsonapi"
	"github.com/mfcochauxlaberge/karigo"
)

// Test ...
func Test(src karigo.Source) error {
	for _, scenario := range scenarios {
		err := src.Reset()
		if err != nil {
			return err
		}

		tx, err := src.Begin()
		if err != nil {
			return err
		}

		for _, step := range scenario.steps {
			switch s := step.(type) {
			case karigo.Op:
				err := tx.Apply([]karigo.Op{s})
				if err != nil {
					return err
				}
			case []karigo.Op:
				err := tx.Apply(s)
				if err != nil {
					return err
				}
			case string:
				switch s {
				case "rollback":
					err = tx.Rollback()
					if err != nil {
						return err
					}
				case "commit":
					err = tx.Commit()
					if err != nil {
						return err
					}
				}
			default:
				return errors.New("karigo: unknown step")
			}
		}

		for set := range scenario.verif {
			col, err := src.Collection(karigo.QueryCol{
				Set: set,
			})
			if err != nil {
				return err
			}
			if col == nil {
				return errors.New("no collection returned")
			}
		}
	}

	return nil

	// Scenario 1

	// ops := []Op{}
	// ops = append(ops, OpAddSet("users")...)
	// ops = append(ops, OpAddAttr("users", "username", "string", false)...)
	// ops = append(ops, OpAddAttr("users", "password", "string", false)...)
	// ops = append(ops, OpAddAttr("users", "created-at", "time", false)...)
	// ops = append(ops, OpAddAttr("users", "updated-at", "time", false)...)

	// ops = append(ops, OpAddSet("articles")...)
	// ops = append(ops, OpAddAttr("articles", "title", "string", false)...)
	// ops = append(ops, OpAddAttr("articles", "content", "string", false)...)
	// ops = append(ops, OpAddAttr("articles", "created-at", "string", false)...)
	// ops = append(ops, OpAddAttr("articles", "updated-at", "string", false)...)

	// ops = append(ops, OpAddRel("users", "articles", false)...)

	// err = tx.Apply(ops)
	// if err != nil {
	// 	return err
	// }

	// err = tx.Commit()
	// if err != nil {
	// 	return err
	// }
}

var scenarios = []scenario{
	scenario{
		steps: []interface{}{
			karigo.OpAddSet("users"),
			karigo.OpAddAttr("users", "username", "string", false),
			karigo.OpAddAttr("users", "password", "string", false),
			karigo.OpAddAttr("users", "created-at", "time", false),
			karigo.OpAddAttr("users", "updated-at", "time", false),
			karigo.OpAddSet("articles"),
			karigo.OpAddAttr("articles", "title", "string", false),
			karigo.OpAddAttr("articles", "content", "string", false),
			karigo.OpAddAttr("articles", "created-at", "string", false),
			karigo.OpAddAttr("articles", "updated-at", "string", false),
			karigo.OpAddRel("users", "articles", false),
		},
		verif: map[string][]jsonapi.Resource{},
	},
}

type scenario struct {
	steps []interface{}
	verif map[string][]jsonapi.Resource
}
