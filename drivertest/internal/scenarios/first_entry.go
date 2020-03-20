package scenarios

import "github.com/mfcochauxlaberge/karigo/query"

func init() {
	Scenarios = append(Scenarios,
		Scenario{
			Name: "first_entry",
			Steps: []interface{}{
				query.NewOpCreateSet("type_name"),
				query.NewOpActivateSet("type_name"),
			},
		},
	)
}
