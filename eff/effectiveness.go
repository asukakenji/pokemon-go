package eff

import (
	"strings"

	"github.com/asukakenji/pokemon-go/lang"
	"github.com/asukakenji/pokemon-go/typ"

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
)

// Effectiveness
type Effectiveness int8

// The "E" prefix is required to make the symbols public.
// Go requires the first character of a public identifier to be in uppercase.
const (
	EX Effectiveness = iota // No Effect (×)
	EA                      // Not Very Effective (▲)
	E_                      // Normal
	EO                      // Super Effective (●)
)

//go:generate stringer -type=Effectiveness

// Official Table: http://www.pokemon.co.jp/whats/summary/
var table = [18][18]Effectiveness{
	/* ATK -> 無  炎  水  草   電  氷  闘  毒   地  飛   超  虫  岩   霊  竜  悪   鋼  妖 */
	/* 無 */ {E_, E_, E_, E_, E_, E_, EO, E_, E_, E_, E_, E_, E_, EX, E_, E_, E_, E_},
	/* 炎 */ {E_, EA, EO, EA, E_, EA, E_, E_, EO, E_, E_, EA, EO, E_, E_, E_, EA, EA},
	/* 水 */ {E_, EA, EA, EO, EO, EA, E_, E_, E_, E_, E_, E_, E_, E_, E_, E_, EA, E_},
	/* 草 */ {E_, EO, EA, EA, EA, EO, E_, EO, EA, EO, E_, EO, E_, E_, E_, E_, E_, E_},
	/* 電 */ {E_, E_, E_, E_, EA, E_, E_, E_, EO, EA, E_, E_, E_, E_, E_, E_, EA, E_},
	/* 氷 */ {E_, EO, E_, E_, E_, EA, EO, E_, E_, E_, E_, E_, EO, E_, E_, E_, EO, E_},
	/* 闘 */ {E_, E_, E_, E_, E_, E_, E_, E_, E_, EO, EO, EA, EA, E_, E_, EA, E_, EO},
	/* 毒 */ {E_, E_, E_, EA, E_, E_, EA, EA, EO, E_, EO, EA, E_, E_, E_, E_, E_, EA},
	/* 地 */ {E_, E_, EO, EO, EX, EO, E_, EA, E_, E_, E_, E_, EA, E_, E_, E_, E_, E_},
	/* 飛 */ {E_, E_, E_, EA, EO, EO, EA, E_, EX, E_, E_, EA, EO, E_, E_, E_, E_, E_},
	/* 超 */ {E_, E_, E_, E_, E_, E_, EA, E_, E_, E_, EA, EO, E_, EO, E_, EO, E_, E_},
	/* 虫 */ {E_, EO, E_, EA, E_, E_, EA, E_, EA, EO, E_, E_, EO, E_, E_, E_, E_, E_},
	/* 岩 */ {EA, EA, EO, EO, E_, E_, EO, EA, EO, EA, E_, E_, E_, E_, E_, E_, EO, E_},
	/* 霊 */ {EX, E_, E_, E_, E_, E_, EX, EA, E_, E_, E_, EA, E_, EO, E_, EO, E_, E_},
	/* 竜 */ {E_, EA, EA, EA, EA, EO, E_, E_, E_, E_, E_, E_, E_, E_, EO, E_, E_, EO},
	/* 悪 */ {E_, E_, E_, E_, E_, E_, EO, E_, E_, E_, EX, EO, E_, EA, E_, EA, E_, EO},
	/* 鋼 */ {EA, EO, E_, EA, E_, EA, EO, EX, EO, EA, EA, EA, EA, E_, EA, E_, EA, EA},
	/* 妖 */ {E_, E_, E_, E_, E_, E_, EA, EO, E_, E_, E_, EA, E_, E_, EX, EA, EO, E_},
}

// moveType: The type of the move (attacker);
// pokemonType: The type of the pokemon (defender).
func For(moveType, pokemonType typ.Type) Effectiveness {
	return table[pokemonType][moveType]
}

func ForEach(f func(Effectiveness)) {
	for i := EX; i <= EO; i = i + 1 {
		f(Effectiveness(i))
	}
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
