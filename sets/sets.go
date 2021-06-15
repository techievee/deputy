package sets

// IntegerSet the data structure to store SETS of uint64, not thread safe with current implementation
type IntegerSet struct {
	items map[uint64]struct{}
}

// Add adds a new element to the Set. Returns a pointer to the Set.
func (s *IntegerSet) Add(t uint64) *IntegerSet {
	if s.items == nil {
		s.items = make(map[uint64]struct{})
	}
	_, ok := s.items[t]
	if !ok {
		s.items[t] = struct{}{}
	}
	return s
}

// Clear removes all elements from the Set
func (s *IntegerSet) Clear() {
	s.items = make(map[uint64]struct{})
}

// Delete removes the int from the Set and returns bool
func (s *IntegerSet) Delete(item uint64) bool {
	_, ok := s.items[item]
	if ok {
		delete(s.items, item)
	}
	return ok
}

// Has returns true if the Set contains the int
func (s *IntegerSet) Has(item uint64) bool {
	_, ok := s.items[item]
	return ok
}

// Items returns the uint64(s) stored
func (s *IntegerSet) Items() []uint64 {
	items := []uint64{}
	for i := range s.items {
		items = append(items, i)
	}
	return items
}

// Size returns the size of the set
func (s *IntegerSet) Size() int {
	return len(s.items)
}
