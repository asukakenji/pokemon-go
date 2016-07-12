package pokemon

import (
	"sort"

	"github.com/asukakenji/pokemon-go/typ"
)

type PokemonSlice []Pokemon

func All() PokemonSlice {
	result := make([]Pokemon, 151)
	for i, pokemon := 0, Bulbasaur; pokemon <= Mew; i, pokemon = i+1, pokemon+1 {
		result[i] = pokemon
	}
	return result
}

func (ps PokemonSlice) ForEach(consumer func(Pokemon)) {
	for _, p := range ps {
		consumer(p)
	}
}

func (ps PokemonSlice) Filter(predicate func(Pokemon) bool) PokemonSlice {
	result := make(PokemonSlice, 0)
	for _, p := range ps {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

func ByTypeF(t typ.Type) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return p.Type1() == t || p.Type2() == t
	}
}

func ByTypesF(t1, t2 typ.Type) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return p.Type1() == t1 && p.Type2() == t2 || p.Type1() == t2 && p.Type2() == t1
	}
}

func (ps PokemonSlice) Sort(less func(Pokemon, Pokemon) bool) PokemonSlice {
	result := make(PokemonSlice, len(ps))
	copy(result, ps)
	sort.Stable(pokemonSlice{result, less})
	return result
}

func ById(p1, p2 Pokemon) bool {
	return p1.Id() < p2.Id()
}

func ByBaseCombatPower(p1, p2 Pokemon) bool {
	return p1.BaseCombatPower() < p2.BaseCombatPower()
}

func ByBaseHitPoints(p1, p2 Pokemon) bool {
	return p1.BaseHitPoints() < p2.BaseHitPoints()
}

// Helper function
func minType(t1, t2 typ.Type) typ.Type {
	if t1 < t2 {
		return t1
	}
	return t2
}

func ByType(p1, p2 Pokemon) bool {
	return minType(p1.Type1(), p1.Type2()) < minType(p2.Type1(), p2.Type2())
}

func ByWeight(p1, p2 Pokemon) bool {
	return p1.Weight() < p2.Weight()
}

func ByHeight(p1, p2 Pokemon) bool {
	return p1.Height() < p2.Height()
}

// Helper type
type pokemonSlice struct {
	slice PokemonSlice
	less  func(Pokemon, Pokemon) bool
}

func (ps pokemonSlice) Len() int {
	return len(ps.slice)
}

func (ps pokemonSlice) Less(i, j int) bool {
	return ps.less(ps.slice[i], ps.slice[j])
}

func (ps pokemonSlice) Swap(i, j int) {
	ps.slice[i], ps.slice[j] = ps.slice[j], ps.slice[i]
}
