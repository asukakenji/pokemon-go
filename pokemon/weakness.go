package pokemon

import (
	"github.com/asukakenji/pokemon-go/typ"
)

type Weakness struct {
	Type       typ.Type // Type
	Multiplier float64  // Multiplier
}

// WeaknessSlice is a helper type for sorting the elements in a Weakness slice.
type WeaknessSlice []Weakness

func (ws WeaknessSlice) Len() int {
	return len(ws)
}

func (ws WeaknessSlice) Less(i, j int) bool {
	return ws[i].Multiplier < ws[j].Multiplier
}

func (ws WeaknessSlice) Swap(i, j int) {
	ws[i], ws[j] = ws[j], ws[i]
}
