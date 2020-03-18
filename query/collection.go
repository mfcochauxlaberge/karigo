package query

import "github.com/mfcochauxlaberge/jsonapi"

// Col ...
type Col struct {
	Set             string
	Fields          []string
	BelongsToFilter jsonapi.BelongsToFilter
	// Filter          *jsonapi.Condition
	Sort       []string
	PageSize   uint
	PageNumber uint
}

// NewCol creates a new Col object from a *jsonapi.URL object.
func NewCol(url *jsonapi.URL) Col {
	var fields []string
	if f, ok := url.Params.Fields[url.ResType]; ok {
		fields = make([]string, len(f))
		copy(fields, f)
	} else {
		fields = []string{"id"}
	}

	query := Col{
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
