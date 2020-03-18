package query

import "github.com/mfcochauxlaberge/jsonapi"

// Res ...
type Res struct {
	Set    string
	ID     string
	Fields []string
}

// NewRes creates a new Res object from a *jsonapi.URL object.
func NewRes(url *jsonapi.URL) Res {
	var fields []string
	if f, ok := url.Params.Fields[url.ResType]; ok {
		fields = make([]string, len(f))
		copy(fields, f)
	} else {
		fields = []string{"id"}
	}

	query := Res{
		Set:    url.ResType,
		ID:     url.ResID,
		Fields: fields,
	}

	return query
}
