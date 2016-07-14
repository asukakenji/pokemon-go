package weak

import (
	"github.com/asukakenji/pokemon-go/generic"
)

// ByMultiplier is designed to be used with Iterable.Sort.
// It returns a function that is usable as a parameter for Sort.
// When ordering is Ascending,
// the returned function returns true if w1 is less than w2 by multiplier.
func ByMultiplier(ordering generic.Ordering) func(Weakness, Weakness) bool {
	if ordering == generic.Ascending {
		return func(w1, w2 Weakness) bool {
			return w1.Multiplier < w2.Multiplier
		}
	}
	return func(w1, w2 Weakness) bool {
		return w1.Multiplier > w2.Multiplier
	}
}
