package pokemon

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
