package weak

import (
	typ "github.com/asukakenji/pokemon-go/type"
)

type Weakness struct {
	Type       typ.Type // Type
	Multiplier float64  // Multiplier
}
