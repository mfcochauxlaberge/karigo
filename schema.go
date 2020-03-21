package karigo

import (
	"github.com/mfcochauxlaberge/jsonapi"
	"github.com/mfcochauxlaberge/karigo/query"
)

// FirstSchema ...
func FirstSchema() *jsonapi.Schema {
	schema := &jsonapi.Schema{}

	typ, err := jsonapi.BuildType(set{})
	if err != nil {
		panic(err)
	}

	err = schema.AddType(typ)
	if err != nil {
		panic(err)
	}

	typ, err = jsonapi.BuildType(attr{})
	if err != nil {
		panic(err)
	}

	err = schema.AddType(typ)
	if err != nil {
		panic(err)
	}

	typ, err = jsonapi.BuildType(rel{})
	if err != nil {
		panic(err)
	}

	err = schema.AddType(typ)
	if err != nil {
		panic(err)
	}

	typ, err = jsonapi.BuildType(meta{})
	if err != nil {
		panic(err)
	}

	err = schema.AddType(typ)
	if err != nil {
		panic(err)
	}

	return schema
}

// set ...
type set struct {
	ID string `json:"id" api:"0_sets"`

	// Attributes
	Name    string `json:"name" api:"attr"`
	Version uint   `json:"version" api:"attr"`
	Created bool   `json:"created" api:"attr"`
	Active  bool   `json:"active" api:"attr"`

	// Relationships
	Attrs []string `json:"attrs" api:"rel,0_attrs,set"`
	Rels  []string `json:"rels" api:"rel,0_rels,from-set"`
}

// attr ...
type attr struct {
	ID string `json:"id" api:"0_attrs"`

	// Attributes
	Name    string `json:"name" api:"attr"`
	Type    string `json:"type" api:"attr"`
	Null    bool   `json:"null" api:"attr"`
	Created bool   `json:"created" api:"attr"`
	Active  bool   `json:"active" api:"attr"`

	// Relationships
	Set string `json:"set" api:"rel,0_sets,attrs"`
}

// rel ...
type rel struct {
	ID string `json:"id" api:"0_rels"`

	// Attributes
	FromName string `json:"from-name" api:"attr"`
	ToOne    bool   `json:"to-one" api:"attr"`
	ToName   string `json:"to-name" api:"attr"`
	FromOne  bool   `json:"from-one" api:"attr"`
	Created  bool   `json:"created" api:"attr"`
	Active   bool   `json:"active" api:"attr"`

	// Relationships
	FromSet string `json:"from-set" api:"rel,0_sets,rels"`
	ToSet   string `json:"to-set" api:"rel,0_sets"`
}

// meta ...
type meta struct {
	ID string `json:"id" api:"0_meta"`

	// Attributes
	Value string `json:"value" api:"attr"`
}

// updateSchema updates the given schema according to the operations.
func updateSchema(schema *jsonapi.Schema, tx query.Tx) error {
	newSchema := &jsonapi.Schema{}

	// Sets
	sets, err := tx.Collection(query.Col{
		Set:      "0_sets",
		PageSize: 999,
	})
	if err != nil {
		return err
	}

	for i := 0; i < sets.Len(); i++ {
		set := sets.At(i)

		err = newSchema.AddType(jsonapi.Type{
			Name: set.Get("name").(string),
		})
		if err != nil {
			return err
		}

		// Attributes
		attrs, err := tx.Collection(query.Col{
			Set: "0_attrs",
			BelongsToFilter: jsonapi.BelongsToFilter{
				Type: "0_sets",
				ID:   set.GetID(),
			},
			PageSize: 999,
		})
		if err != nil {
			return err
		}

		for j := 0; j < attrs.Len(); j++ {
			attr := attrs.At(i)

			t, _ := jsonapi.GetAttrType(attr.Get("type").(string))

			err = newSchema.AddAttr(
				set.GetID(),
				jsonapi.Attr{
					Name:     attr.Get("name").(string),
					Type:     t,
					Nullable: attr.Get("nullable").(bool),
				},
			)
			if err != nil {
				return err
			}
		}

		// Relationships
		rels, err := tx.Collection(query.Col{
			Set: "0_rels",
			BelongsToFilter: jsonapi.BelongsToFilter{
				Type: "0_sets",
				ID:   set.GetID(),
			},
			PageSize: 999,
		})
		if err != nil {
			return err
		}

		for j := 0; j < rels.Len(); j++ {
			rel := rels.At(i)

			err = newSchema.AddRel(
				set.GetID(),
				jsonapi.Rel{
					FromType: rel.GetToOne("from-set"),
					FromName: rel.Get("from-name").(string),
					ToOne:    rel.Get("to-one").(bool),
					ToType:   rel.GetToOne("to-set"),
					ToName:   rel.Get("to-name").(string),
					FromOne:  rel.Get("from-one").(bool),
				},
			)
			if err != nil {
				return err
			}
		}
	}

	errs := newSchema.Check()
	if len(errs) > 0 {
		return errs[0]
	}

	*schema = *newSchema

	return nil
}

// func activateSet(s *jsonapi.Schema, name string) error {
// 	err := s.AddType(jsonapi.Type{
// 		Name: name,
// 	})

// 	return err
// }

// func deactivateSet(s *jsonapi.Schema, res jsonapi.Resource) {
// 	s.RemoveType(res.GetID())
// }

// func activateAttr(s *jsonapi.Schema, res jsonapi.Resource) error {
// 	typ, null := jsonapi.GetAttrType(res.Get("type").(string))
// 	err := s.AddAttr(res.GetToOne("set"), jsonapi.Attr{
// 		Name:     res.Get("name").(string),
// 		Type:     typ,
// 		Nullable: null,
// 	})

// 	return err
// }

// func deactivateAttr(s *jsonapi.Schema, res jsonapi.Resource) {
// 	s.RemoveAttr(res.GetToOne("set"), res.Get("Name").(string))
// }

// func activateRel(s *jsonapi.Schema, res jsonapi.Resource) error {
// 	rel := jsonapi.Rel{
// 		FromType: res.GetToOne("from-set"),
// 		FromName: res.Get("from-name").(string),
// 		ToOne:    res.Get("to-one").(bool),
// 		ToType:   res.GetToOne("to-set"),
// 		ToName:   res.Get("to-name").(string),
// 		FromOne:  res.Get("from-one").(bool),
// 	}
// 	rel = rel.Normalize()

// 	err := s.AddRel(res.GetToOne("from-set"), rel)
// 	if err != nil {
// 		return err
// 	}

// 	err = s.AddRel(res.GetToOne("to-set"), rel.Invert())
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func deactivateRel(s *jsonapi.Schema, res jsonapi.Resource) {
// 	s.RemoveRel(res.Get("from-type").(string), res.Get("name").(string))
// }
