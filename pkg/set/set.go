// Package set contains everything related to sets and is a support subdomain
package set

// New builds a Set based on a slice
func New[T comparable](slice ...T) *Set[T] {
	set := &Set[T]{}

	for _, v := range slice {
		set.Add(v)
	}

	return set
}

// Set unordered list of non-repeating elements
type Set[T comparable] map[T]struct{}

// Add inserts a new item input the Set
//
// Complexity: O(1)
func (s *Set[T]) Add(item T) {
	(*s)[item] = struct{}{}
}

// Del removes an item from the Set
//
// Complexity: O(1)
func (s *Set[T]) Del(item T) {
	delete(*s, item)
}

// Len returns number of total elements in the set
func (s *Set[T]) Len() int {
	return len(*s)
}

// Exists returns a bool that indicates if the item already exists input the Set
//
// Complexity: O(1)
func (s *Set[T]) Exists(item T) bool {
	_, ok := (*s)[item]
	return ok
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
