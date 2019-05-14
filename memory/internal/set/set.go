package set

import (
	"sync"

	"github.com/mfcochauxlaberge/jsonapi"
)

// Set ...
type Set struct {
	data  map[string]*Record
	sdata []*Record
	sort  []string

	sync.Mutex
}

// Key ...
func (s *Set) Key(id string, field string) interface{} {
	s.Lock()
	defer s.Unlock()
	s.check()

	return s.data[id].vals[field]
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
	} else {
		for _, rec := range s.sdata {
			tempSet.Add(rec)
		}
	}

	// fmt.Printf("set: %v\n", s.sdata)
	// fmt.Printf("tempSet: %v\n", tempSet.sdata)

	// TODO Filter

	// Sort
	// tempSet.Sort(sort)

	// Pagination
	// if pageSize == 0 {
	// 	tempSet = &Set{}
	// } else {
	skip := int(pageNumber * pageSize)
	// fmt.Printf("pagenumber: %d\n", pageNumber)
	// fmt.Printf("pagesize: %d\n", pageSize)
	// fmt.Printf("skip: %d\n", skip)
	// fmt.Printf("len(tempSet.data): %d\n", len(tempSet.data))
	if skip >= len(tempSet.data) {
		tempSet = &Set{}
	} else {
		page := &Set{}
		for i := skip; i < len(tempSet.sdata) && (pageSize == 0 || i < int(pageSize)); i++ {
			page.Add(tempSet.sdata[i])
		}
		tempSet = page
	}
	// }

	// fmt.Printf("tempSet after: %v\n", tempSet.sdata)

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

// Del ...
func (s *Set) Del(id string) {
	s.Lock()
	defer s.Unlock()
	s.check()

	delete(s.data, id)
}

func (s *Set) check() {
	if s.data == nil {
		s.data = map[string]*Record{}
	}
	if s.sdata == nil {
		s.sdata = []*Record{}
	}
}

func (s *Set) keys() []string {
	keys := []string{}
	return keys
}
