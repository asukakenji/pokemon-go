package weak

import (
	"sort"

	"github.com/asukakenji/pokemon-go/generic"
)

// Iterable defines an interface for an iterable collection of Weakness.
type Iterable interface {
	// ForEach iterates through and apply the consumer function to each element
	// in the Iterable.
	ForEach(consumer func(Weakness))
	// Filter iterates through and apply the predicate function to each element
	// in the Iterable.
	// Elements resulting true are collected and returned as another Iterable.
	Filter(predicate func(Weakness) bool) Iterable
	// Map iterates through and apply the mapper function to each element
	// in the Iterable.
	// The return values are collected and returned as a generic.Iterable.
	Map(mapper func(Weakness) interface{}) generic.Iterable
	// Sort sorts (stably) the elements in the Iterable.
	// The result is returned as another Iterable.
	Sort(less func(Weakness, Weakness) bool) Iterable
}

// Slice implements the Iterable interface.
// It represents a slice of Weakness.
type Slice []Weakness

// ForEach implements the same method in the Iterable interface.
func (s Slice) ForEach(consumer func(Weakness)) {
	for _, e := range s {
		consumer(e)
	}
}

// Filter implements the same method in the Iterable interface.
func (s Slice) Filter(predicate func(Weakness) bool) Iterable {
	result := make(Slice, 0)
	for _, e := range s {
		if predicate(e) {
			result = append(result, e)
		}
	}
	return result
}

// Map implements the same method in the Iterable interface.
func (s Slice) Map(mapper func(Weakness) interface{}) generic.Iterable {
	result := make(generic.Slice, 0)
	for _, e := range s {
		result = append(result, mapper(e))
	}
	return result
}

// Sort implements the same method in the Iterable interface.
func (s Slice) Sort(less func(Weakness, Weakness) bool) Iterable {
	result := make(Slice, len(s))
	copy(result, s)
	sort.Stable(sortableSlice{result, less})
	return result
}

// TODO: Write Doc!
func Wrap(iterable generic.Iterable) Iterable {
	return wrappedIterable{iterable}
}

// TODO: Write Doc!
type wrappedIterable struct {
	iterable generic.Iterable
}

// ForEach implements the same method in the Iterable interface.
func (s wrappedIterable) ForEach(consumer func(Weakness)) {
	s.iterable.ForEach(func(e interface{}) {
		consumer(e.(Weakness))
	})
}

// Filter implements the same method in the Iterable interface.
func (s wrappedIterable) Filter(predicate func(Weakness) bool) Iterable {
	result := make(Slice, 0)
	s.iterable.ForEach(func(e interface{}) {
		e2 := e.(Weakness)
		if predicate(e2) {
			result = append(result, e2)
		}
	})
	return result
}

// Map implements the same method in the Iterable interface.
func (s wrappedIterable) Map(mapper func(Weakness) interface{}) generic.Iterable {
	result := make(generic.Slice, 0)
	s.iterable.ForEach(func(e interface{}) {
		result = append(result, mapper(e.(Weakness)))
	})
	return result
}

// Sort implements the same method in the Iterable interface.
func (s wrappedIterable) Sort(less func(Weakness, Weakness) bool) Iterable {
	result := make(Slice, 0)
	s.iterable.ForEach(func(e interface{}) {
		result = append(result, e.(Weakness))
	})
	sort.Stable(sortableSlice{result, less})
	return result
}

// sortableSlice implements the sort.Interface interface.
// It serves as an adapter for sort.Interface,
// and as a wrapper for a Slice and a custom comparison function.
type sortableSlice struct {
	slice Slice
	less  func(Weakness, Weakness) bool
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
