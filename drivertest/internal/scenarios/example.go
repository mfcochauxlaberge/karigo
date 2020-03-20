package scenarios

import "github.com/mfcochauxlaberge/karigo/query"

func init() {
	Scenarios = append(Scenarios,
		Scenario{
			Name: "example",
			Steps: []interface{}{
				// Users
				query.NewOpCreateSet("users"),
				query.NewOpActivateSet("users"),
				query.NewOpCreateAttr("users", "username", "string", false),
				query.NewOpActivateAttr("users", "username"),
				query.NewOpCreateAttr("users", "password", "string", false),
				query.NewOpActivateAttr("users", "password"),
				query.NewOpCreateAttr("users", "created-at", "time.Time", false),
				query.NewOpActivateAttr("users", "created-at"),

				// Articles
				query.NewOpCreateSet("articles"),
				query.NewOpActivateSet("articles"),
				query.NewOpCreateAttr("articles", "title", "string", false),
				query.NewOpActivateAttr("articles", "title"),
				query.NewOpCreateAttr("articles", "content", "string", false),
				query.NewOpActivateAttr("articles", "content"),
				query.NewOpCreateAttr("articles", "created-at", "string", false),
				query.NewOpActivateAttr("articles", "created-at"),
				query.NewOpCreateAttr("articles", "updated-at", "string", false),
				query.NewOpActivateAttr("articles", "updated-at"),

				// Relationships
				query.NewOpCreateRel(
					"users",
					"articles",
					"articles",
					"author",
					false,
					true,
				),
				query.NewOpActivateRel("articles_author_users_articles"),

				// Data
				query.NewOpSet("users", "", "id", "abc123"),
				query.NewOpSet("users", "abc123", "username", "mafiaboy"),
			},
		},
	)
}
