package pokemon

import (
	"strings"

	"github.com/asukakenji/pokemon-go/lang"

	"github.com/asukakenji/pokemon-go/internal/effectiveness/de"
	"github.com/asukakenji/pokemon-go/internal/effectiveness/en"
	"github.com/asukakenji/pokemon-go/internal/effectiveness/es"
	"github.com/asukakenji/pokemon-go/internal/effectiveness/fr"
	"github.com/asukakenji/pokemon-go/internal/effectiveness/it"
	"github.com/asukakenji/pokemon-go/internal/effectiveness/ja"
	"github.com/asukakenji/pokemon-go/internal/effectiveness/ko"
	"github.com/asukakenji/pokemon-go/internal/effectiveness/zh-CN"
	"github.com/asukakenji/pokemon-go/internal/effectiveness/zh-HK"
	"github.com/asukakenji/pokemon-go/internal/effectiveness/zh-TW"
)

// Effectiveness
type Effectiveness int8

// The "E" prefix is required to make the symbols public.
// Go requires the first character of a public identifier to be in uppercase.
const (
	EX Effectiveness = iota // No Effect (×)
	EA                      // Not Very Effective (▲)
	E_                      // Normal ( )
	EO                      // Super Effective (●)
)

//go:generate stringer -type=Effectiveness

// moveType: The type of the move (attacker);
// pokemonType: The type of the pokemon (defender).
func EffectivenessFor(moveType, pokemonType Type) Effectiveness {
	return table[pokemonType][moveType]
}

func (e Effectiveness) Id() int {
	return int(e)
}

func (e Effectiveness) LocalName(l lang.Language) string {
	switch l {
	case lang.Japanese:
		return ja.Effectiveness(e).String()
	case lang.English:
		return strings.Replace(en.Effectiveness(e).String(), "_", " ", -1)
	case lang.French:
		return fr.Effectiveness(e).String()
	case lang.German:
		return de.Effectiveness(e).String()
	case lang.Italian:
		return it.Effectiveness(e).String()
	case lang.Korean:
		return ko.Effectiveness(e).String()
	case lang.Spanish:
		return es.Effectiveness(e).String()
	case lang.ChineseChina:
		return zhCN.Effectiveness(e).String()
	case lang.ChineseHongKong:
		return zhHK.Effectiveness(e).String()
	case lang.ChineseTaiwan:
		return zhTW.Effectiveness(e).String()
	default:
		return ""
	}
}

const _Effectiveness_Symbol_string = "×▲ ●"

var _Effectiveness_Symbol_index = [...]uint8{0, 2, 5, 6, 9}

func (e Effectiveness) Symbol() string {
	if e < 0 || e >= Effectiveness(len(_Effectiveness_Symbol_index)-1) {
		panic("Invalid Effectiveness")
	}
	return _Effectiveness_Symbol_string[_Effectiveness_Symbol_index[e]:_Effectiveness_Symbol_index[e+1]]
}

func (e Effectiveness) Multiplier() float64 {
	return [4]float64{0.0, 0.5, 1.0, 2.0}[e.Id()]
}
