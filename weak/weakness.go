package weak

import (
	typ "github.com/asukakenji/pokemon-go/type"
)

type Weakness struct {
	Type       typ.Type // Type
	Multiplier float64  // Multiplier
}

// ja: 弱点
// en: Weakness(es)
// fr: Faiblesse(s)
// de: Schwäche(n)
// it: Debolezza (Debolezze)
// es: Debilidad
// zh-CN: 弱点
// zh-HK: 弱點
// zh-TW: 弱點
