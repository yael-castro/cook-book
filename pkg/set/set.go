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

// Put inserts a new item input the Set
//
// Complexity: O(1)
func (s *Set[T]) Put(item T) {
	(*s)[item] = struct{}{}
}

// Has returns a bool that indicates if the item already exists input the Set
//
// Complexity: O(1)
func (s *Set[T]) Has(item T) bool {
	_, ok := (*s)[item]
	return ok
}

// Remove removes an item from the Set
//
// Complexity: O(1)
func (s *Set[T]) Remove(item T) {
	delete(*s, item)
}

// Slice transform the Set input a slice
//
// Complexity: O(n)
//
// n - is the Set length
func (s *Set[T]) Slice() []T {
	slice := make([]T, 0, len(*s))

	for k := range *s {
		slice = append(slice, k)
	}

	return slice
}
