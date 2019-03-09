package scenarios

import (
	"github.com/mfcochauxlaberge/jsonapi"
	"github.com/mfcochauxlaberge/karigo"
)

func init() {
	Scenarios = append(Scenarios,
		Scenario{
			Steps: []interface{}{
				karigo.OpAddSet("users"),
				karigo.OpAddAttr("users", "username", "string", false),
				karigo.OpAddAttr("users", "password", "string", false),
				karigo.OpAddAttr("users", "created-at", "time", false),
				karigo.OpAddSet("articles"),
				karigo.OpAddAttr("articles", "title", "string", false),
				karigo.OpAddAttr("articles", "content", "string", false),
				karigo.OpAddAttr("articles", "created-at", "string", false),
				karigo.OpAddAttr("articles", "updated-at", "string", false),
				karigo.OpAddRel("users", "articles", false),
			},
			Verif: map[string][]jsonapi.Resource{},
		},
	)
}
