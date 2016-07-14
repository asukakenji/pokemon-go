package weak

import (
	"github.com/asukakenji/pokemon-go/generic"
)

// TODO: Write Doc!
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
