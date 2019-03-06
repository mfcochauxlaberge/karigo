package karigo

import "github.com/kkaribu/jsonapi"

// NewSchema ...
func NewSchema() *Schema {
	return &Schema{
		schema:      []jsonapi.Type{},
		getFuncs:    map[string]Tx{},
		createFuncs: map[string]Tx{},
		updateFuncs: map[string]Tx{},
		deleteFuncs: map[string]Tx{},
	}
}

// Schema ...
type Schema struct {
	schema      []jsonapi.Type
	getFuncs    map[string]Tx
	createFuncs map[string]Tx
	updateFuncs map[string]Tx
	deleteFuncs map[string]Tx
}
