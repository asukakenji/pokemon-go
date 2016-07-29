package weak

import (
	"github.com/asukakenji/pokemon-go/lang"
	typ "github.com/asukakenji/pokemon-go/type"
)

type Weakness struct {
	Type       typ.Type // Type
	Multiplier float64  // Multiplier
}

func PackageName(l lang.Language) string {
	switch l {
	case lang.Japanese:
		return "弱点"
	case lang.English:
		return "Weakness(es)"
	case lang.French:
		return "Faiblesse(s)"
	case lang.German:
		return "Schwäche(n)"
	case lang.Italian:
		return "Debolezz(a/e)"
	case lang.Korean:
		return "약점"
	case lang.Spanish:
		return "Debilidad"
	case lang.ChineseSimplified:
		return "弱点"
	case lang.ChineseTraditional:
		return "弱點"
	case lang.ChineseChina:
		return "弱点"
	case lang.ChineseHongKong:
		return "弱點"
	case lang.ChineseTaiwan:
		return "弱點"
	default:
		return ""
	}
}
