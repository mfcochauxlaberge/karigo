package scenarios

import "github.com/mfcochauxlaberge/karigo"

func init() {
	Scenarios = append(Scenarios,
		Scenario{
			Name: "first_entry",
			Steps: []interface{}{
				karigo.NewOpCreateSet("type_name"),
				karigo.NewOpActivateSet("type_name"),
			},
		},
	)
}
