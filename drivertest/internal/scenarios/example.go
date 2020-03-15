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
				karigo.NewOpCreateSet("users"),
				karigo.NewOpActivateSet("users"),
				karigo.NewOpCreateAttr("users", "username", "string", false),
				karigo.NewOpActivateAttr("users", "username"),
				karigo.NewOpCreateAttr("users", "password", "string", false),
				karigo.NewOpActivateAttr("users", "password"),
				karigo.NewOpCreateAttr("users", "created-at", "time.Time", false),
				karigo.NewOpActivateAttr("users", "created-at"),

				// Articles
				karigo.NewOpCreateSet("articles"),
				karigo.NewOpActivateSet("articles"),
				karigo.NewOpCreateAttr("articles", "title", "string", false),
				karigo.NewOpActivateAttr("articles", "title"),
				karigo.NewOpCreateAttr("articles", "content", "string", false),
				karigo.NewOpActivateAttr("articles", "content"),
				karigo.NewOpCreateAttr("articles", "created-at", "string", false),
				karigo.NewOpActivateAttr("articles", "created-at"),
				karigo.NewOpCreateAttr("articles", "updated-at", "string", false),
				karigo.NewOpActivateAttr("articles", "updated-at"),

				// Relationships
				karigo.NewOpCreateRel(
					"users",
					"articles",
					"articles",
					"author",
					false,
					true,
				),
				karigo.NewOpActivateRel("articles_author_users_articles"),

				// Data
				karigo.NewOpSet("users", "", "id", "abc123"),
				karigo.NewOpSet("users", "abc123", "username", "mafiaboy"),
			},
		},
	)
}
