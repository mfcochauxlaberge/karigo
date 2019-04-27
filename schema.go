package karigo

import "github.com/mfcochauxlaberge/jsonapi"

// NewSchema ...
func NewSchema() *Schema {
	return &Schema{
		jaSchema: &jsonapi.Schema{},
		funcs:    map[string]Tx{},
	}
}

// Schema ...
type Schema struct {
	jaSchema *jsonapi.Schema
	funcs    map[string]Tx
}
