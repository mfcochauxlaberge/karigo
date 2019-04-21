package karigo

import "github.com/mfcochauxlaberge/jsonapi"

// NewSchema ...
func NewSchema() *Schema {
	return &Schema{
		jaSchema:    &jsonapi.Schema{},
		getFuncs:    map[string]Tx{},
		createFuncs: map[string]Tx{},
		updateFuncs: map[string]Tx{},
		deleteFuncs: map[string]Tx{},
	}
}

// Schema ...
type Schema struct {
	jaSchema    *jsonapi.Schema
	getFuncs    map[string]Tx
	createFuncs map[string]Tx
	updateFuncs map[string]Tx
	deleteFuncs map[string]Tx
}
