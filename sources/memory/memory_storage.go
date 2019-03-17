package memory

import (
	"github.com/mfcochauxlaberge/jsonapi"
)

type record struct {
	schema *jsonapi.Schema
	id     string
	vals   map[string]interface{}
}
