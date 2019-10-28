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
				karigo.NewOpActivateAttr("users", "username"),
				karigo.NewOpAddAttr("users", "password", "string", false),
				karigo.NewOpActivateAttr("users", "password"),
				karigo.NewOpAddAttr("users", "created-at", "time", false),
				karigo.NewOpActivateAttr("users", "created-at"),

				// Articles
				karigo.NewOpAddSet("articles"),
				karigo.NewOpActivateSet("articles"),
				karigo.NewOpAddAttr("articles", "title", "string", false),
				karigo.NewOpActivateAttr("articles", "title"),
				karigo.NewOpAddAttr("articles", "content", "string", false),
				karigo.NewOpActivateAttr("articles", "content"),
				karigo.NewOpAddAttr("articles", "created-at", "string", false),
				karigo.NewOpActivateAttr("articles", "created-at"),
				karigo.NewOpAddAttr("articles", "updated-at", "string", false),
				karigo.NewOpActivateAttr("articles", "updated-at"),

				// Relationships
				karigo.NewOpAddRel(
					"users",
					"articles",
					"articles",
					"author",
					false,
					true,
				),
				karigo.NewOpActivateRel("users", "articles"),

				// Data
				karigo.NewOpSet("users", "", "id", "abc123"),
				karigo.NewOpSet("users", "abc123", "username", "mafiaboy"),
			},
		},
	)
}
