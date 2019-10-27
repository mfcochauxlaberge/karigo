package scenarios

import (
	"github.com/mfcochauxlaberge/karigo"
)

func init() {
	Scenarios = append(Scenarios,
		Scenario{
			Name: "example",
			Steps: []interface{}{
				// Users
				karigo.NewOpAddSet("users"),
				karigo.NewOpActivateSet("users"),
				karigo.NewOpAddAttr("users", "username", "string", false),
				karigo.NewOpAddAttr("users", "password", "string", false),
				karigo.NewOpAddAttr("users", "created-at", "time", false),

				// Articles
				karigo.NewOpAddSet("articles"),
				karigo.NewOpActivateSet("articles"),
				karigo.NewOpAddAttr("articles", "title", "string", false),
				karigo.NewOpAddAttr("articles", "content", "string", false),
				karigo.NewOpAddAttr("articles", "created-at", "string", false),
				karigo.NewOpAddAttr("articles", "updated-at", "string", false),

				// Relationships
				karigo.NewOpAddRel(
					"users",
					"articles",
					"articles",
					"author",
					false,
					true,
				),

				// Data
				karigo.NewOpSet("users", "", "id", "abc123"),
			},
		},
	)
}
