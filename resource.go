package karigo

import (
	"github.com/mfcochauxlaberge/jsonapi"
)

// NewResource ...
func NewResource() Resource {
	return Resource{}
}

// Resource ...
type Resource struct {
	Type jsonapi.Type

	vals map[string]interface{}
}

// Attrs ...
func (r *Resource) Attrs() []jsonapi.Attr {
	// TODO Fix ordering
	attrs := make([]jsonapi.Attr, 0, len(r.Type.Attrs))
	for n := range r.Type.Attrs {
		attrs = append(attrs, r.Type.Attrs[n])
	}
	return attrs
}

// Rels ...
func (r *Resource) Rels() []jsonapi.Rel {
	// TODO Fix ordering
	rels := make([]jsonapi.Rel, 0, len(r.Type.Rels))
	for n := range r.Type.Rels {
		rels = append(rels, r.Type.Rels[n])
	}
	return rels
}

// Attr ...
func (r *Resource) Attr(key string) jsonapi.Attr {
	if attr, ok := r.Type.Attrs[key]; ok {
		return attr
	}
	return jsonapi.Attr{}
}

// Rel ...
func (r *Resource) Rel(key string) jsonapi.Rel {
	if rel, ok := r.Type.Rels[key]; ok {
		return rel
	}
	return jsonapi.Rel{}
}

// New ...
func (r *Resource) New() Resource {
	return Resource{}
}

// GetID ...
func (r *Resource) GetID() string {
	return ""
}

// GetType ...
func (r *Resource) GetType() string {
	return r.Type.Name
}

// Get ...
func (r *Resource) Get(key string) interface{} {
	return nil
}

// SetID ...
func (r *Resource) SetID(id string) {}

// Set ...
func (r *Resource) Set(key string, val interface{}) {}

// GetToOne ...
func (r *Resource) GetToOne(key string) string {
	return ""
}

// GetToMany ...
func (r *Resource) GetToMany(key string) []string {
	return nil
}

// SetToOne ...
func (r *Resource) SetToOne(key string, rel string) {}

// SetToMany ...
func (r *Resource) SetToMany(key string, rels []string) {}

// Validate ...
func (r *Resource) Validate() []error {
	return nil
}

// Copy ...
func (r *Resource) Copy() Resource {
	return *r
}

// UnmarshalJSON ...
func (r *Resource) UnmarshalJSON(payload []byte) error {
	return nil
}
