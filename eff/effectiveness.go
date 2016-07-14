package eff

import (
	"strings"

	"github.com/asukakenji/pokemon-go/eff/internal/de"
	"github.com/asukakenji/pokemon-go/eff/internal/en"
	"github.com/asukakenji/pokemon-go/eff/internal/es"
	"github.com/asukakenji/pokemon-go/eff/internal/fr"
	"github.com/asukakenji/pokemon-go/eff/internal/it"
	"github.com/asukakenji/pokemon-go/eff/internal/ja"
	"github.com/asukakenji/pokemon-go/eff/internal/ko"
	"github.com/asukakenji/pokemon-go/eff/internal/zh-CN"
	"github.com/asukakenji/pokemon-go/eff/internal/zh-HK"
	"github.com/asukakenji/pokemon-go/eff/internal/zh-TW"

	"github.com/asukakenji/pokemon-go/lang"
	typ "github.com/asukakenji/pokemon-go/type"
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
func For(moveType, pokemonType typ.Type) Effectiveness {
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

const _Symbol_string = "×▲ ●"

var _Symbol_index = [...]uint8{0, 2, 5, 6, 9}

func (e Effectiveness) Symbol() string {
	if e < 0 || e >= Effectiveness(len(_Symbol_index)-1) {
		panic("Invalid Effectiveness")
	}
	return _Symbol_string[_Symbol_index[e]:_Symbol_index[e+1]]
}

func (e Effectiveness) Multiplier() float64 {
	return [4]float64{0.0, 0.5, 1.0, 2.0}[e.Id()]
}