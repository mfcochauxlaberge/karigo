package scenarios

import (
	"github.com/mfcochauxlaberge/jsonapi"
	"github.com/mfcochauxlaberge/karigo"
)

func init() {
	Scenarios = append(Scenarios,
		Scenario{
			Steps: []interface{}{
				karigo.NewOpAddSet("type_name"),
			},
			Verif: map[string][]jsonapi.Resource{},
		},
	)
}
