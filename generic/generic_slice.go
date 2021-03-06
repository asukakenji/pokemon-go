package generic

import (
	"sort"
)

// Iterable defines an interface for an iterable collection of interface{}.
type Iterable interface {
	// ForEach iterates through and apply the consumer function to each element
	// in the Iterable.
	ForEach(consumer func(interface{}))
	// Filter iterates through and apply the predicate function to each element
	// in the Iterable.
	// Elements resulting true are collected and returned as another Iterable.
	Filter(predicate func(interface{}) bool) Iterable
	// Map iterates through and apply the mapper function to each element
	// in the Iterable.
	// The return values are collected and returned as another Iterable.
	Map(mapper func(interface{}) interface{}) Iterable
	// Sort sorts (stably) the elements in the Iterable.
	// The result is returned as another Iterable.
	Sort(less func(interface{}, interface{}) bool) Iterable
}

// Slice implements the Iterable interface.
// It represents a slice of interface{}.
type Slice []interface{}

// ForEach implements the same method in the Iterable interface.
func (s Slice) ForEach(consumer func(interface{})) {
	for _, e := range s {
		consumer(e)
	}
}

// Filter implements the same method in the Iterable interface.
func (s Slice) Filter(predicate func(interface{}) bool) Iterable {
	result := make(Slice, 0)
	for _, e := range s {
		if predicate(e) {
			result = append(result, e)
		}
	}
	return result
}

// Map implements the same method in the Iterable interface.
func (s Slice) Map(mapper func(interface{}) interface{}) Iterable {
	result := make(Slice, 0)
	for _, e := range s {
		result = append(result, mapper(e))
	}
	return result
}

// Sort implements the same method in the Iterable interface.
func (s Slice) Sort(less func(interface{}, interface{}) bool) Iterable {
	result := make(Slice, len(s))
	copy(result, s)
	sort.Stable(sortableSlice{result, less})
	return result
}

// sortableSlice implements the sort.Interface interface.
// It serves as an adapter for sort.Interface,
// and as a wrapper for a Slice and a custom comparison function.
type sortableSlice struct {
	slice Slice
	less  func(interface{}, interface{}) bool
}

// Len implements the same method in the sort.Interface interface.
func (s sortableSlice) Len() int {
	return len(s.slice)
}

// Less implements the same method in the sort.Interface interface.
func (s sortableSlice) Less(i, j int) bool {
	return s.less(s.slice[i], s.slice[j])
}

// Swap implements the same method in the sort.Interface interface.
func (s sortableSlice) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}
