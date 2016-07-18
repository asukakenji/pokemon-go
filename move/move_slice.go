package move

import (
	"sort"

	"github.com/asukakenji/pokemon-go/generic"
)

// Iterable defines an interface for an iterable collection of Move.
type Iterable interface {
	// ForEach iterates through and apply the consumer function to each element
	// in the Iterable.
	ForEach(consumer func(Move))
	// Filter iterates through and apply the predicate function to each element
	// in the Iterable.
	// Elements resulting true are collected and returned as another Iterable.
	Filter(predicate func(Move) bool) Iterable
	// Map iterates through and apply the mapper function to each element
	// in the Iterable.
	// The return values are collected and returned as a generic.Iterable.
	Map(mapper func(Move) interface{}) generic.Iterable
	// Sort sorts (stably) the elements in the Iterable.
	// The result is returned as another Iterable.
	Sort(less func(Move, Move) bool) Iterable
}

// Slice implements the Iterable interface.
// It represents a slice of Move.
type Slice []Move

// ForEach implements the same method in the Iterable interface.
func (s Slice) ForEach(consumer func(Move)) {
	for _, e := range s {
		consumer(e)
	}
}

// Filter implements the same method in the Iterable interface.
func (s Slice) Filter(predicate func(Move) bool) Iterable {
	result := make(Slice, 0)
	for _, e := range s {
		if predicate(e) {
			result = append(result, e)
		}
	}
	return result
}

// Map implements the same method in the Iterable interface.
func (s Slice) Map(mapper func(Move) interface{}) generic.Iterable {
	result := make(generic.Slice, 0)
	for _, e := range s {
		result = append(result, mapper(e))
	}
	return result
}

// Sort implements the same method in the Iterable interface.
func (s Slice) Sort(less func(Move, Move) bool) Iterable {
	result := make(Slice, len(s))
	copy(result, s)
	sort.Stable(sortableSlice{result, less})
	return result
}

// All returns all meaningful values of Move as an Iterable.
func All() Iterable {
	return virtualSlice(0)
}

// virtualSlice implements the Iterable interface.
// It is the concrete type of the result returned by All().
// Instead of building a slice of all values,
// the elements are generated on-the-fly when needed.
type virtualSlice int

// ForEach implements the same method in the Iterable interface.
func (s virtualSlice) ForEach(consumer func(Move)) {
	for e := Acid; e <= ZenHeadbutt; e++ {
		consumer(e)
	}
	for e := AerialAce; e <= X_Scissor; e++ {
		consumer(e)
	}
}

// Filter implements the same method in the Iterable interface.
func (s virtualSlice) Filter(predicate func(Move) bool) Iterable {
	result := make(Slice, 0)
	for e := Acid; e <= ZenHeadbutt; e++ {
		if predicate(e) {
			result = append(result, e)
		}
	}
	for e := AerialAce; e <= X_Scissor; e++ {
		if predicate(e) {
			result = append(result, e)
		}
	}
	return result
}

// Map implements the same method in the Iterable interface.
func (s virtualSlice) Map(mapper func(Move) interface{}) generic.Iterable {
	result := make(generic.Slice, 0)
	for e := Acid; e <= ZenHeadbutt; e++ {
		result = append(result, mapper(e))
	}
	for e := AerialAce; e <= X_Scissor; e++ {
		result = append(result, mapper(e))
	}
	return result
}

// Sort implements the same method in the Iterable interface.
func (s virtualSlice) Sort(less func(Move, Move) bool) Iterable {
	result := make(Slice, int16(X_Scissor)-int16(Acid)+1)
	i := 0
	for e := Acid; e <= ZenHeadbutt; e++ {
		result[i] = e
		i++
	}
	for e := AerialAce; e <= X_Scissor; e++ {
		result[i] = e
		i++
	}
	sort.Stable(sortableSlice{result, less})
	return result
}

// sortableSlice implements the sort.Interface interface.
// It serves as an adapter for sort.Interface,
// and as a wrapper for a Slice and a custom comparison function.
type sortableSlice struct {
	slice Slice
	less  func(Move, Move) bool
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
