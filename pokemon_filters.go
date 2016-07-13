package pokemon

// ByTypeF is designed to be used with PokemonIterable.Filter. It returns
// a function that is usable as a parameter for Filter. The returned function
// returns true when one of the types of the Pokemon matches the query.
func ByTypeF(t Type) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return p.Type1() == t || p.Type2() == t
	}
}

// ByTypesF is designed to be used with PokemonIterable.Filter. It returns
// a function that is usable as a parameter for Filter. The returned function
// returns true when both of the types of the Pokemon match the query, in any
// order.
func ByTypesF(t1, t2 Type) func(Pokemon) bool {
	return func(p Pokemon) bool {
		return p.Type1() == t1 && p.Type2() == t2 || p.Type1() == t2 && p.Type2() == t1
	}
}
