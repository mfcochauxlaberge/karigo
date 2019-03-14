package scenarios

import (
	"github.com/mfcochauxlaberge/jsonapi"
	"github.com/mfcochauxlaberge/karigo"
)

func init() {
	Scenarios = append(Scenarios,
		Scenario{
			Steps: []interface{}{
				// Users
				karigo.NewOpAddSet("users"),
				karigo.NewOpAddAttr("users", "username", "string", false),
				karigo.NewOpAddAttr("users", "password", "string", false),
				karigo.NewOpAddAttr("users", "created-at", "time", false),

				// Articles
				karigo.NewOpAddSet("articles"),
				karigo.NewOpAddAttr("articles", "title", "string", false),
				karigo.NewOpAddAttr("articles", "content", "string", false),
				karigo.NewOpAddAttr("articles", "created-at", "string", false),
				karigo.NewOpAddAttr("articles", "updated-at", "string", false),

				// Relationships
				karigo.NewOpAddRel("users", "articles", false),

				// Data
				karigo.NewOpSet("users", "", "id", "abc123"),
			},
			Verif: map[string][]jsonapi.Resource{},
		},
	)
}
