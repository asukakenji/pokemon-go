package pokemon

import (
	"sort"

	"github.com/asukakenji/pokemon-go/typ"
)

// PokemonIterable defines an interface for an iterable collection of Pokemon.
type PokemonIterable interface {
	// ForEach iterates through and apply the consumer function to each element
	// in the PokemonIterable.
	ForEach(consumer func(Pokemon))
	// Filter iterates through and apply the predicate function to each element
	// in the PokemonIterable. Elements returning true are returned as another
	// PokemonIterable.
	Filter(predicate func(Pokemon) bool) PokemonIterable
	// Sort sorts (stably) the elements in the PokemonIterable. The result is
	// returned as another PokemonIterable.
	Sort(less func(Pokemon, Pokemon) bool) PokemonIterable
}

// PokemonSlice implements the PokemonIterable interface. It represents a slice
// of Pokemon.
type PokemonSlice []Pokemon

// ForEach implements the same method in the PokemonIterable interface.
func (ps PokemonSlice) ForEach(consumer func(Pokemon)) {
	for _, p := range ps {
		consumer(p)
	}
}

// Filter implements the same method in the PokemonIterable interface.
func (ps PokemonSlice) Filter(predicate func(Pokemon) bool) PokemonIterable {
	result := make(PokemonSlice, 0)
	for _, p := range ps {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

// Sort implements the same method in the PokemonIterable interface.
func (ps PokemonSlice) Sort(less func(Pokemon, Pokemon) bool) PokemonIterable {
	result := make(PokemonSlice, len(ps))
	copy(result, ps)
	sort.Stable(sortablePokemonSlice{result, less})
	return result
}

// AllPokemons returns all Pokemon's as a PokemonIterable
func AllPokemons() PokemonIterable {
	return virtualPokemonSlice(0)
}

// ByTypeF is designed to be used with PokemonIterable.Filter. It returns
// a function that is usable as a parameter for Filter. The returned function
// returns true when one of the types of the Pokemon matches the query.
func ByTypeF(t typ.Type) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return p.Type1() == t || p.Type2() == t
	}
}

// ByTypesF is designed to be used with PokemonIterable.Filter. It returns
// a function that is usable as a parameter for Filter. The returned function
// returns true when both of the types of the Pokemon match the query, in any
// order.
func ByTypesF(t1, t2 typ.Type) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return p.Type1() == t1 && p.Type2() == t2 || p.Type1() == t2 && p.Type2() == t1
	}
}

// ById is designed to be used with PokemonIterable.Sort. It returns
// true if p1 is less than p2 by id.
func ById(p1, p2 Pokemon) bool {
	return p1.Id() < p2.Id()
}

// ByBaseCombatPower is designed to be used with PokemonIterable.Sort. It returns
// true if p1 is less than p2 by base combat power (base CP).
func ByBaseCombatPower(p1, p2 Pokemon) bool {
	return p1.BaseCombatPower() < p2.BaseCombatPower()
}

// ByBaseHitPoints is designed to be used with PokemonIterable.Sort. It returns
// true if p1 is less than p2 by base hit points (base HP).
func ByBaseHitPoints(p1, p2 Pokemon) bool {
	return p1.BaseHitPoints() < p2.BaseHitPoints()
}

// ByWeight is designed to be used with PokemonIterable.Sort. It returns
// true if p1 is less than p2 by weight.
func ByWeight(p1, p2 Pokemon) bool {
	return p1.Weight() < p2.Weight()
}

// ByHeight is designed to be used with PokemonIterable.Sort. It returns
// true if p1 is less than p2 by height.
func ByHeight(p1, p2 Pokemon) bool {
	return p1.Height() < p2.Height()
}

// virtualPokemonSlice implements the PokemonIterable interface. It is the
// concrete type of the result returned by All(). The elements are generated as
// needed.
type virtualPokemonSlice int

// ForEach implements the same method in the PokemonIterable interface.
func (ps virtualPokemonSlice) ForEach(consumer func(Pokemon)) {
	for p := Bulbasaur; p <= Mew; p++ {
		consumer(p)
	}
}

// Filter implements the same method in the PokemonIterable interface.
func (ps virtualPokemonSlice) Filter(predicate func(Pokemon) bool) PokemonIterable {
	result := make(PokemonSlice, 0)
	for p := Bulbasaur; p <= Mew; p++ {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

// Sort implements the same method in the PokemonIterable interface.
func (ps virtualPokemonSlice) Sort(less func(Pokemon, Pokemon) bool) PokemonIterable {
	result := make(PokemonSlice, 151)
	for i, p := 0, Bulbasaur; p <= Mew; i, p = i+1, p+1 {
		result[i] = p
	}
	sort.Stable(sortablePokemonSlice{result, less})
	return result
}

// sortablePokemonSlice implements the sort.Interface interface. It serves as
// an adapter for sort.Interface and a wrapper for a PokemonSlice and a custom
// comparison function.
type sortablePokemonSlice struct {
	slice PokemonSlice
	less  func(Pokemon, Pokemon) bool
}

// Len implements the same method in the sort.Interface interface.
func (ps sortablePokemonSlice) Len() int {
	return len(ps.slice)
}

// Less implements the same method in the sort.Interface interface.
func (ps sortablePokemonSlice) Less(i, j int) bool {
	return ps.less(ps.slice[i], ps.slice[j])
}

// Swap implements the same method in the sort.Interface interface.
func (ps sortablePokemonSlice) Swap(i, j int) {
	ps.slice[i], ps.slice[j] = ps.slice[j], ps.slice[i]
}
