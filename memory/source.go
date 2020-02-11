package memory

import (
	"sync"

	"github.com/mfcochauxlaberge/karigo"

	"github.com/mfcochauxlaberge/jsonapi"
)

// Source ...
type Source struct {
	sets map[string]*jsonapi.SoftCollection

	sync.Mutex
}

// Connect ...
func (s *Source) Connect(_ map[string]string) error {
	return nil
}

// Ping ...
func (s *Source) Ping() bool {
	return true
}

// Reset ...
func (s *Source) Reset() error {
	s.Lock()
	defer s.Unlock()

	types := map[string]*jsonapi.Type{}

	for _, typ := range karigo.FirstSchema().Types {
		ctyp := typ.Copy()
		types[ctyp.Name] = &ctyp
	}

	s.sets = map[string]*jsonapi.SoftCollection{}

	// Sets
	for _, typ := range types {
		s.sets[typ.Name] = &jsonapi.SoftCollection{}
		s.sets[typ.Name].SetType(typ)
	}

	// Types and attributes
	for _, typ := range types {
		typ := typ.Copy()

		attrIDs := []string{}
		relIDs := []string{}

		for _, field := range typ.Fields() {
			if attr, ok := typ.Attrs[field]; ok {
				attrIDs = append(attrIDs, typ.Name+"_"+attr.Name)
			} else if rel, ok := typ.Rels[field]; ok {
				if rel.FromType+rel.FromName ==
					rel.Invert().FromType+rel.Invert().FromName {
					relIDs = append(relIDs, rel.String())
				}
			}
		}

		s.sets["0_sets"].Add(makeSoftResource(
			types["0_sets"],
			typ.Name,
			map[string]interface{}{
				"name":    typ.Name,
				"version": 0,
				"created": true,
				"active":  true,
				"attrs":   attrIDs,
				"rels":    relIDs,
			},
		))

		// 0_attrs
		for _, attr := range typ.Attrs {
			s.sets["0_attrs"].Add(makeSoftResource(
				types["0_attrs"],
				typ.Name+"_"+attr.Name,
				map[string]interface{}{
					"name":    attr.Name,
					"type":    jsonapi.GetAttrTypeString(attr.Type, false),
					"null":    attr.Nullable,
					"created": true,
					"active":  true,
					"set":     typ.Name,
				},
			))
		}
	}

	// Relationships
	for _, rel := range karigo.FirstSchema().Rels() {
		s.sets["0_rels"].Add(makeSoftResource(
			types["0_rels"],
			rel.String(),
			map[string]interface{}{
				"from-name": rel.FromName,
				"to-one":    rel.ToOne,
				"to-name":   rel.ToName,
				"from-one":  rel.FromOne,
				"created":   true,
				"active":    true,
				"from-set":  rel.FromType,
				"to-set":    rel.ToType,
			},
		))
	}

	return nil
}

// NewTx ...
func (s *Source) NewTx() (karigo.Tx, error) {
	return &Tx{
		src:  s,
		sets: s.sets,
	}, nil
}
