package set

// Exists is the default map value
var Exists = struct{}{}

// Set is the main interface
type Set struct {
	m map[interface{}]struct{}
}

// New creates and returns a new set
func New(items ...interface{}) *Set {
	s := &Set{}
	s.m = make(map[interface{}]struct{})
	s.Add(items...)

	return s
}

// Add inserts element in set. Multi elements can be added at the same time, and duplicate will be ignored.
func (s *Set) Add(items ...interface{}) error {
	for _, item := range items {
		s.m[item] = Exists
	}

	return nil
}

// Size returns how many elements are in set
func (s *Set) Size() int {
	return len(s.m)
}

// Contains checks if a given item is in set, return true or false
func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

// Clear removes all elements in set, make it as new
func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}

// Copy returns a new set that have exactly the same elements
func (s *Set) Copy() *Set {
	newSet := New()
	for key := range s.m {
		newSet.Add(key)
	}

	return newSet
}

// Remove deletes element from set. Multiple elements can be removed at once,
// and non-exists element will be ignoed without exception
func (s *Set) Remove(items ...interface{}) {
	for _, item := range items {
		delete(s.m, item)
	}
}

// Equal checks if two sets contain exactly the same elements
func (s *Set) Equal(other *Set) bool {
	if s.Size() != other.Size() {
		return false
	}

	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// IsSubset checks if all elements are in other set
func (s *Set) IsSubset(other *Set) bool {
	if s.Size() > other.Size() {
		return false
	}

	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// IsSuperset checks if all elements in other set are in set
func (s *Set) IsSuperset(other *Set) bool {
	return other.IsSubset(s)
}

// Union returns a set containing all elements that are in s or other set
func (s *Set) Union(other *Set) *Set {
	newSet := other.Copy()

	for key := range s.m {
		newSet.Add(key)
	}
	return newSet
}

// Intersection returns a new set containing all elements that are in both sets
func (s *Set) Intersection(other *Set) *Set {
	newSet := New()

	for key := range s.m {
		if other.Contains(key) {
			newSet.Add(key)
		}
	}
	return newSet
}

// Difference returns a new set containing elements only in original set
func (s *Set) Difference(other *Set) *Set {
	newSet := s.Copy()

	for key := range s.m {
		if other.Contains(key) {
			newSet.Remove(key)
		}
	}
	return newSet
}

// SymmetricDifference returna a new set containing elements that only in one set not the other
func (s *Set) SymmetricDifference(other *Set) *Set {
	newSet := s.Difference(other).Union(other.Difference(s))

	return newSet
}
