package pokemon

import (
	"github.com/asukakenji/pokemon-go/generic"
)

// ById is designed to be used with Iterable.Sort.
// It returns a function that is usable as a parameter for Sort.
// When ordering is Ascending,
// the returned function returns true if p1 is less than p2 by id.
func ById(ordering generic.Ordering) func(Pokemon, Pokemon) bool {
	if ordering == generic.Ascending {
		return func(p1, p2 Pokemon) bool {
			return p1.Id() < p2.Id()
		}
	}
	return func(p1, p2 Pokemon) bool {
		return p1.Id() < p2.Id()
	}
}

// ByBaseCombatPower is designed to be used with Iterable.Sort.
// It returns a function that is usable as a parameter for Sort.
// When ordering is Ascending,
// the returned function returns true if p1 if p1 is less than p2 by base combat
// power (base CP).
func ByBaseCombatPower(ordering generic.Ordering) func(Pokemon, Pokemon) bool {
	if ordering == generic.Ascending {
		return func(p1, p2 Pokemon) bool {
			return p1.BaseCombatPower() < p2.BaseCombatPower()
		}
	}
	return func(p1, p2 Pokemon) bool {
		return p1.BaseCombatPower() < p2.BaseCombatPower()
	}
}

// ByBaseHitPoints is designed to be used with Iterable.Sort.
// It returns a function that is usable as a parameter for Sort.
// When ordering is Ascending,
// the returned function returns true if p1 if p1 is less than p2 by base hit
// points (base HP).
func ByBaseHitPoints(ordering generic.Ordering) func(Pokemon, Pokemon) bool {
	if ordering == generic.Ascending {
		return func(p1, p2 Pokemon) bool {
			return p1.BaseHitPoints() < p2.BaseHitPoints()
		}
	}
	return func(p1, p2 Pokemon) bool {
		return p1.BaseHitPoints() < p2.BaseHitPoints()
	}
}

// ByWeight is designed to be used with Iterable.Sort.
// It returns a function that is usable as a parameter for Sort.
// When ordering is Ascending,
// the returned function returns true if p1 is less than p2 by weight.
func ByWeight(ordering generic.Ordering) func(Pokemon, Pokemon) bool {
	if ordering == generic.Ascending {
		return func(p1, p2 Pokemon) bool {
			return p1.Weight() < p2.Weight()
		}
	}
	return func(p1, p2 Pokemon) bool {
		return p1.Weight() < p2.Weight()
	}
}

// ByHeight is designed to be used with Iterable.Sort.
// It returns a function that is usable as a parameter for Sort.
// When ordering is Ascending,
// the returned function returns true if p1 is less than p2 by height.
func ByHeight(ordering generic.Ordering) func(Pokemon, Pokemon) bool {
	if ordering == generic.Ascending {
		return func(p1, p2 Pokemon) bool {
			return p1.Height() < p2.Height()
		}
	}
	return func(p1, p2 Pokemon) bool {
		return p1.Height() < p2.Height()
	}
}
