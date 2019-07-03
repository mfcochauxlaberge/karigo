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
			Verif: []string{
				"0_attrs.0_attrs_active.active=true",
				"0_attrs.0_attrs_active.name=active",
				"0_attrs.0_attrs_active.null=false",
				"0_attrs.0_attrs_active.set=0_attrs",
				"0_attrs.0_attrs_active.type=bool",
				"0_attrs.0_attrs_name.active=true",
				"0_attrs.0_attrs_name.name=name",
				"0_attrs.0_attrs_name.null=false",
				"0_attrs.0_attrs_name.set=0_attrs",
				"0_attrs.0_attrs_name.type=string",
				"0_attrs.0_attrs_null.active=true",
				"0_attrs.0_attrs_null.name=null",
				"0_attrs.0_attrs_null.null=false",
				"0_attrs.0_attrs_null.set=0_attrs",
				"0_attrs.0_attrs_null.type=bool",
				"0_attrs.0_attrs_type.active=true",
				"0_attrs.0_attrs_type.name=type",
				"0_attrs.0_attrs_type.null=false",
				"0_attrs.0_attrs_type.set=0_attrs",
				"0_attrs.0_attrs_type.type=string",
				"0_attrs.0_meta_value.active=true",
				"0_attrs.0_meta_value.name=value",
				"0_attrs.0_meta_value.null=false",
				"0_attrs.0_meta_value.set=0_meta",
				"0_attrs.0_meta_value.type=string",
				"0_attrs.0_rels_active.active=true",
				"0_attrs.0_rels_active.name=active",
				"0_attrs.0_rels_active.null=false",
				"0_attrs.0_rels_active.set=0_rels",
				"0_attrs.0_rels_active.type=bool",
				"0_attrs.0_rels_name.active=true",
				"0_attrs.0_rels_name.name=name",
				"0_attrs.0_rels_name.null=false",
				"0_attrs.0_rels_name.set=0_rels",
				"0_attrs.0_rels_name.type=string",
				"0_attrs.0_rels_to-one.active=true",
				"0_attrs.0_rels_to-one.name=to-one",
				"0_attrs.0_rels_to-one.null=false",
				"0_attrs.0_rels_to-one.set=0_rels",
				"0_attrs.0_rels_to-one.type=bool",
				"0_attrs.0_sets_active.active=true",
				"0_attrs.0_sets_active.name=active",
				"0_attrs.0_sets_active.null=false",
				"0_attrs.0_sets_active.set=0_sets",
				"0_attrs.0_sets_active.type=bool",
				"0_attrs.0_sets_name.active=true",
				"0_attrs.0_sets_name.name=name",
				"0_attrs.0_sets_name.null=false",
				"0_attrs.0_sets_name.set=0_sets",
				"0_attrs.0_sets_name.type=string",
				"0_attrs.0_sets_version.active=true",
				"0_attrs.0_sets_version.name=version",
				"0_attrs.0_sets_version.null=false",
				"0_attrs.0_sets_version.set=0_sets",
				"0_attrs.0_sets_version.type=int",
				"0_attrs.articles_content.active=true",
				"0_attrs.articles_content.name=content",
				"0_attrs.articles_content.null=false",
				"0_attrs.articles_content.set=articles",
				"0_attrs.articles_content.type=string",
				"0_attrs.articles_created-at.active=true",
				"0_attrs.articles_created-at.name=created-at",
				"0_attrs.articles_created-at.null=false",
				"0_attrs.articles_created-at.set=articles",
				"0_attrs.articles_created-at.type=string",
				"0_attrs.articles_title.active=true",
				"0_attrs.articles_title.name=title",
				"0_attrs.articles_title.null=false",
				"0_attrs.articles_title.set=articles",
				"0_attrs.articles_title.type=string",
				"0_attrs.articles_updated-at.active=true",
				"0_attrs.articles_updated-at.name=updated-at",
				"0_attrs.articles_updated-at.null=false",
				"0_attrs.articles_updated-at.set=articles",
				"0_attrs.articles_updated-at.type=string",
				"0_attrs.users_created-at.active=true",
				"0_attrs.users_created-at.name=created-at",
				"0_attrs.users_created-at.null=false",
				"0_attrs.users_created-at.set=users",
				"0_attrs.users_created-at.type=time",
				"0_attrs.users_password.active=true",
				"0_attrs.users_password.name=password",
				"0_attrs.users_password.null=false",
				"0_attrs.users_password.set=users",
				"0_attrs.users_password.type=string",
				"0_attrs.users_username.active=true",
				"0_attrs.users_username.name=username",
				"0_attrs.users_username.null=false",
				"0_attrs.users_username.set=users",
				"0_attrs.users_username.type=string",
				"0_rels.0_attrs_set.active=true",
				"0_rels.0_attrs_set.inverse=0_sets_attrs",
				"0_rels.0_attrs_set.name=set",
				"0_rels.0_attrs_set.set=0_attrs",
				"0_rels.0_attrs_set.to-one=true",
				"0_rels.0_rels_inverse.active=true",
				"0_rels.0_rels_inverse.inverse=0_rels_inverse",
				"0_rels.0_rels_inverse.name=inverse",
				"0_rels.0_rels_inverse.set=0_rels",
				"0_rels.0_rels_inverse.to-one=true",
				"0_rels.0_rels_set.active=true",
				"0_rels.0_rels_set.inverse=0_sets_rels",
				"0_rels.0_rels_set.name=set",
				"0_rels.0_rels_set.set=0_rels",
				"0_rels.0_rels_set.to-one=true",
				"0_rels.0_sets_attrs.active=true",
				"0_rels.0_sets_attrs.inverse=0_attrs_set",
				"0_rels.0_sets_attrs.name=attrs",
				"0_rels.0_sets_attrs.set=0_sets",
				"0_rels.0_sets_attrs.to-one=false",
				"0_rels.0_sets_rels.active=true",
				"0_rels.0_sets_rels.inverse=0_rels_set",
				"0_rels.0_sets_rels.name=rels",
				"0_rels.0_sets_rels.set=0_sets",
				"0_rels.0_sets_rels.to-one=false",
				"0_rels.users_articles.active=true",
				"0_rels.users_articles.inverse=",
				"0_rels.users_articles.name=articles",
				"0_rels.users_articles.set=users",
				"0_rels.users_articles.to-one=false",
				"0_sets.0_attrs.active=true",
				"0_sets.0_attrs.attrs=0_attrs_name,0_attrs_type,0_attrs_null,0_attrs_active",
				"0_sets.0_attrs.name=0_attrs",
				"0_sets.0_attrs.rels=0_attrs_set",
				"0_sets.0_attrs.version=0",
				"0_sets.0_meta.active=true",
				"0_sets.0_meta.attrs=0_meta_value",
				"0_sets.0_meta.name=0_meta",
				"0_sets.0_meta.rels=",
				"0_sets.0_meta.version=0",
				"0_sets.0_rels.active=true",
				"0_sets.0_rels.attrs=0_rels_name,0_rels_to-one,0_rels_active",
				"0_sets.0_rels.name=0_rels",
				"0_sets.0_rels.rels=0_rels_set",
				"0_sets.0_rels.version=0",
				"0_sets.0_sets.active=true",
				"0_sets.0_sets.attrs=0_sets_name,0_sets_version,0_sets_active",
				"0_sets.0_sets.name=0_sets",
				"0_sets.0_sets.rels=0_sets_attrs,0_sets_rels",
				"0_sets.0_sets.version=0",
				"0_sets.articles.active=true",
				"0_sets.articles.attrs=",
				"0_sets.articles.name=articles",
				"0_sets.articles.rels=",
				"0_sets.articles.version=0",
				"0_sets.users.active=true",
				"0_sets.users.attrs=",
				"0_sets.users.name=users",
				"0_sets.users.rels=",
				"0_sets.users.version=0",
				"users.abc123.users_articles=",
				"users.abc123.users_password=",
				"users.abc123.users_username=",
			},
		},
	)
}
