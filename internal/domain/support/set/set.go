// Package set contains everything related to sets and is a support subdomain
package set

// New builds a Set based on a slice
func New[T comparable](slice []T) *Set[T] {
	set := &Set[T]{}

	for _, v := range slice {
		set.Put(v)
	}

	return set
}

// Set list of non-repeating elements
type Set[T comparable] map[T]struct{}

// Put inserts a new item in the Set
//
// Complexity: O(1)
func (s *Set[T]) Put(item T) error {
	(*s)[item] = struct{}{}
	return nil
}

// Exists returns a bool that indicates if the item already exists in the Set
//
// Complexity: O(1)
func (s *Set[T]) Exists(item T) bool {
	_, ok := (*s)[item]
	return ok
}

// Remove removes an item from the Set
//
// Complexity: O(1)
func (s *Set[T]) Remove(item T) {
	delete(*s, item)
}

// Slice transform the Set in a slice
//
// Complexity: O(n)
//
// n - is the Set length
func (s *Set[T]) Slice() (slice []T) {
	for k := range *s {
		slice = append(slice, k)
	}

	return
}
