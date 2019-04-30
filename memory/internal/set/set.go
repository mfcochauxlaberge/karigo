package set

import (
	"strings"
	"sync"
	"time"

	"github.com/mfcochauxlaberge/jsonapi"
)

// Set ...
type Set struct {
	data  map[string]*Record
	sdata []*Record
	sort  []string

	sync.Mutex
}

// Resource ...
func (s *Set) Resource(id string, fields []string) jsonapi.Resource {
	s.Lock()
	defer s.Unlock()
	s.check()

	return s.data[id].Resource(fields)
}

// Collection ...
func (s *Set) Collection(ids []string, _ *jsonapi.Condition, sort []string, fields []string, pageSize uint, pageNumber uint) []jsonapi.Resource {
	s.Lock()
	defer s.Unlock()
	s.check()

	tempSet := &Set{}

	// Filter IDs
	if len(ids) > 0 {
		for _, rec := range s.sdata {
			for _, id := range ids {
				if rec.id == id {
					tempSet.Add(rec)
				}
			}
		}
	}

	// TODO Filter

	// Sort
	tempSet.Sort(sort)

	// Pagination
	if pageSize == 0 {
		tempSet = &Set{}
	} else {
		skip := int(pageNumber * pageSize)
		if skip >= len(tempSet.data) {
			tempSet = &Set{}
		} else {
			page := &Set{}
			for i := skip; i < len(tempSet.sdata) || i < int(pageSize); i++ {
				page.Add(tempSet.sdata[i])
			}
			tempSet = page
		}
	}

	// Make the collection
	col := []jsonapi.Resource{}
	for _, rec := range tempSet.sdata {
		col = append(col, rec.Resource(fields))
	}

	return col
}

// Add ...
func (s *Set) Add(rec *Record) {
	s.Lock()
	defer s.Unlock()
	s.check()

	s.data[rec.id] = rec

	// rec already exists
	if _, ok := s.data[rec.id]; ok {
		for i := range s.sdata {
			if s.sdata[i].id == rec.id {
				s.sdata[i] = rec
			}
		}
	}

	// rec does not exist
	s.sdata = append(s.sdata, rec)
}

// Set ...
func (s *Set) Set(id string, field string, v interface{}) {
	s.Lock()
	defer s.Unlock()
	s.check()

	// id should already exist
	if _, ok := s.data[id]; ok {
		s.data[id].vals[field] = v
	}
}

// Sort ...
func (s *Set) Sort(sort []string) {

}

// Len ...
func (s *Set) Len() int { return len(s.data) }

// Swap ...
func (s *Set) Swap(i, j int) { s.sdata[i], s.sdata[j] = s.sdata[j], s.sdata[i] }

// Less ...
func (s *Set) Less(i, j int) bool {
	less := false

	for _, r := range s.sort {
		inverse := false
		if strings.HasPrefix(r, "-") {
			r = r[1:]
			inverse = true
		}

		switch v := s.sdata[i].vals[r].(type) {
		case string:
			if v == s.sdata[j].vals[r].(string) {
				continue
			}
			if inverse {
				return v > s.sdata[j].vals[r].(string)
			}
			return v < s.sdata[j].vals[r].(string)
		case int:
			if v == s.sdata[j].vals[r].(int) {
				continue
			}
			if inverse {
				return v > s.sdata[j].vals[r].(int)
			}
			return v < s.sdata[j].vals[r].(int)
		case int8:
			if v == s.sdata[j].vals[r].(int8) {
				continue
			}
			if inverse {
				return v > s.sdata[j].vals[r].(int8)
			}
			return v < s.sdata[j].vals[r].(int8)
		case int16:
			if v == s.sdata[j].vals[r].(int16) {
				continue
			}
			if inverse {
				return v > s.sdata[j].vals[r].(int16)
			}
			return v < s.sdata[j].vals[r].(int16)
		case int32:
			if v == s.sdata[j].vals[r].(int32) {
				continue
			}
			if inverse {
				return v > s.sdata[j].vals[r].(int32)
			}
			return v < s.sdata[j].vals[r].(int32)
		case int64:
			if v == s.sdata[j].vals[r].(int64) {
				continue
			}
			if inverse {
				return v > s.sdata[j].vals[r].(int64)
			}
			return v < s.sdata[j].vals[r].(int64)
		case uint:
			if v == s.sdata[j].vals[r].(uint) {
				continue
			}
			if inverse {
				return v > s.sdata[j].vals[r].(uint)
			}
			return v < s.sdata[j].vals[r].(uint)
		case uint8:
			if v == s.sdata[j].vals[r].(uint8) {
				continue
			}
			if inverse {
				return v > s.sdata[j].vals[r].(uint8)
			}
			return v < s.sdata[j].vals[r].(uint8)
		case uint16:
			if v == s.sdata[j].vals[r].(uint16) {
				continue
			}
			if inverse {
				return v > s.sdata[j].vals[r].(uint16)
			}
			return v < s.sdata[j].vals[r].(uint16)
		case uint32:
			if v == s.sdata[j].vals[r].(uint32) {
				continue
			}
			if inverse {
				return v > s.sdata[j].vals[r].(uint32)
			}
			return v < s.sdata[j].vals[r].(uint32)
		case bool:
			if v == s.sdata[j].vals[r].(bool) {
				continue
			}
			if inverse {
				return v
			}
			return !v
		case time.Time:
			if v.Equal(s.sdata[j].vals[r].(time.Time)) {
				continue
			}
			if inverse {
				return v.After(s.sdata[j].vals[r].(time.Time))
			}
			return v.Before(s.sdata[j].vals[r].(time.Time))
		case *string:
			if *v == *(s.sdata[j].vals[r].(*string)) {
				continue
			}
			if inverse {
				return *v > *(s.sdata[j].vals[r].(*string))
			}
			return *v < *(s.sdata[j].vals[r].(*string))
		case *int:
			if *v == *(s.sdata[j].vals[r].(*int)) {
				continue
			}
			if inverse {
				return *v > *(s.sdata[j].vals[r].(*int))
			}
			return *v < *(s.sdata[j].vals[r].(*int))
		case *int8:
			if *v == *(s.sdata[j].vals[r].(*int8)) {
				continue
			}
			if inverse {
				return *v > *(s.sdata[j].vals[r].(*int8))
			}
			return *v < *(s.sdata[j].vals[r].(*int8))
		case *int16:
			if *v == *(s.sdata[j].vals[r].(*int16)) {
				continue
			}
			if inverse {
				return *v > *(s.sdata[j].vals[r].(*int16))
			}
			return *v < *(s.sdata[j].vals[r].(*int16))
		case *int32:
			if *v == *(s.sdata[j].vals[r].(*int32)) {
				continue
			}
			if inverse {
				return *v > *(s.sdata[j].vals[r].(*int32))
			}
			return *v < *(s.sdata[j].vals[r].(*int32))
		case *int64:
			if *v == *(s.sdata[j].vals[r].(*int64)) {
				continue
			}
			if inverse {
				return *v > *(s.sdata[j].vals[r].(*int64))
			}
			return *v < *(s.sdata[j].vals[r].(*int64))
		case *uint:
			if *v == *(s.sdata[j].vals[r].(*uint)) {
				continue
			}
			if inverse {
				return *v > *(s.sdata[j].vals[r].(*uint))
			}
			return *v < *(s.sdata[j].vals[r].(*uint))
		case *uint8:
			if *v == *(s.sdata[j].vals[r].(*uint8)) {
				continue
			}
			if inverse {
				return *v > *(s.sdata[j].vals[r].(*uint8))
			}
			return *v < *(s.sdata[j].vals[r].(*uint8))
		case *uint16:
			if *v == *(s.sdata[j].vals[r].(*uint16)) {
				continue
			}
			if inverse {
				return *v > *(s.sdata[j].vals[r].(*uint16))
			}
			return *v < *(s.sdata[j].vals[r].(*uint16))
		case *uint32:
			if *v == *(s.sdata[j].vals[r].(*uint32)) {
				continue
			}
			if inverse {
				return *v > *(s.sdata[j].vals[r].(*uint32))
			}
			return *v < *(s.sdata[j].vals[r].(*uint32))
		case *bool:
			if *v == *(s.sdata[j].vals[r].(*bool)) {
				continue
			}
			if inverse {
				return *v
			}
			return !*v
		case *time.Time:
			if v.Equal(*(s.sdata[j].vals[r].(*time.Time))) {
				continue
			}
			if inverse {
				return v.After(*(s.sdata[j].vals[r].(*time.Time)))
			}
			return v.Before(*(s.sdata[j].vals[r].(*time.Time)))
		}
	}

	return less
}

func (s *Set) check() {
	if s.data == nil {
		s.data = map[string]*Record{}
	}
	if s.sdata == nil {
		s.sdata = []*Record{}
	}
}
