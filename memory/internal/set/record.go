package set

import "github.com/mfcochauxlaberge/jsonapi"

// NewRecord ...
func NewRecord(id string, vals map[string]interface{}) *Record {
	return &Record{
		id:   id,
		vals: vals,
	}
}

// Record ...
type Record struct {
	id   string
	vals map[string]interface{}
}

// Resource ...
func (r *Record) Resource(fields []string) jsonapi.Resource {
	res := &jsonapi.SoftResource{}
	res.SetID(r.id)
	for field := range r.vals {
		for _, f := range fields {
			if field == f {
				res.Set(field, r.vals[field])
				break
			}
		}
	}
	return res
}
