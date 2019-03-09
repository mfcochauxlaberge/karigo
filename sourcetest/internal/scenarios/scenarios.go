package scenarios

import "github.com/mfcochauxlaberge/jsonapi"

// Scenarios ...
var Scenarios = []Scenario{}

// Scenario ...
type Scenario struct {
	Steps []interface{}
	Verif map[string][]jsonapi.Resource
}
