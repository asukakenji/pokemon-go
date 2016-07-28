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

// ByBaseStamina is designed to be used with Iterable.Sort.
// It returns a function that is usable as a parameter for Sort.
// When ordering is Ascending,
// the returned function returns true if p1 if p1 is less than p2
// by base stamina.
func ByBaseStamina(ordering generic.Ordering) func(Pokemon, Pokemon) bool {
	if ordering == generic.Ascending {
		return func(p1, p2 Pokemon) bool {
			return p1.BaseStamina() < p2.BaseStamina()
		}
	}
	return func(p1, p2 Pokemon) bool {
		return p1.BaseStamina() < p2.BaseStamina()
	}
}

// ByBaseAttack is designed to be used with Iterable.Sort.
// It returns a function that is usable as a parameter for Sort.
// When ordering is Ascending,
// the returned function returns true if p1 if p1 is less than p2
// by base attack.
func ByBaseAttack(ordering generic.Ordering) func(Pokemon, Pokemon) bool {
	if ordering == generic.Ascending {
		return func(p1, p2 Pokemon) bool {
			return p1.BaseAttack() < p2.BaseAttack()
		}
	}
	return func(p1, p2 Pokemon) bool {
		return p1.BaseAttack() < p2.BaseAttack()
	}
}

// ByBaseDefense is designed to be used with Iterable.Sort.
// It returns a function that is usable as a parameter for Sort.
// When ordering is Ascending,
// the returned function returns true if p1 if p1 is less than p2
// by base defense.
func ByBaseDefense(ordering generic.Ordering) func(Pokemon, Pokemon) bool {
	if ordering == generic.Ascending {
		return func(p1, p2 Pokemon) bool {
			return p1.BaseDefense() < p2.BaseDefense()
		}
	}
	return func(p1, p2 Pokemon) bool {
		return p1.BaseDefense() < p2.BaseDefense()
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
