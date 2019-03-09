package karigo

import (
	"github.com/mfcochauxlaberge/jsonapi"
)

// QueryCol ...
type QueryCol struct {
	Set             string
	Fields          []string
	BelongsToFilter jsonapi.BelongsToFilter
	// Filter          *jsonapi.Condition
	Sort       []string
	PageSize   int
	PageNumber int
}

// NewQueryCol creates a new *QueryCol object from a *jsonapi.URL object.
func NewQueryCol(url *jsonapi.URL) *QueryCol {
	var fields []string
	if f, ok := url.Params.Fields[url.ResType]; ok {
		fields = make([]string, len(f))
		copy(fields, f)
	} else {
		fields = []string{"id"}
	}

	query := &QueryCol{
		Set:             url.ResType,
		Fields:          fields,
		BelongsToFilter: url.BelongsToFilter,
		// Filter:          url.Params.Filter,
		Sort:       url.Params.SortingRules,
		PageSize:   url.Params.PageSize,
		PageNumber: url.Params.PageNumber,
	}

	return query
}

// QueryRes ...
type QueryRes struct {
	Set    string
	ID     string
	Fields []string
}

// NewQueryRes creates a new *Query object from a *jsonapi.URL object.
func NewQueryRes(url *jsonapi.URL) *QueryRes {
	var fields []string
	if f, ok := url.Params.Fields[url.ResType]; ok {
		fields = make([]string, len(f))
		copy(fields, f)
	} else {
		fields = []string{"id"}
	}

	query := &QueryRes{
		Set:    url.ResType,
		ID:     url.ResID,
		Fields: fields,
	}

	return query
}
