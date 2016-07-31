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
	"github.com/asukakenji/pokemon-go/eff/internal/zh-CHS"
	"github.com/asukakenji/pokemon-go/eff/internal/zh-CHT"
	"github.com/asukakenji/pokemon-go/eff/internal/zh-CN"
	"github.com/asukakenji/pokemon-go/eff/internal/zh-HK"
	"github.com/asukakenji/pokemon-go/eff/internal/zh-TW"

	"github.com/asukakenji/pokemon-go/lang"
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

func PackageName(l lang.Language) string {
	switch l {
	case lang.Japanese:
		return "効果"
	case lang.English:
		return "Effectiveness"
	case lang.French:
		return "Efficacité"
	case lang.German:
		return "Effektivität"
	case lang.Italian:
		return "Efficacia"
	case lang.Korean:
		return "효과"
	case lang.Spanish:
		return "Eficacia"
	case lang.ChineseSimplified:
		return "效果"
	case lang.ChineseTraditional:
		return "效果"
	case lang.ChineseChina:
		return "效果"
	case lang.ChineseHongKong:
		return "效果"
	case lang.ChineseTaiwan:
		return "效果"
	default:
		return ""
	}
}

func (e Effectiveness) Id() int {
	return int(e)
}

func (e Effectiveness) IsValid() bool {
	return EX <= e && e <= EO
}

func (e Effectiveness) LocalName(l lang.Language) string {
	result := ""
	switch l {
	case lang.Japanese:
		result = ja.Effectiveness(e).String()
	case lang.English:
		result = en.Effectiveness(e).String()
	case lang.French:
		result = fr.Effectiveness(e).String()
	case lang.German:
		result = de.Effectiveness(e).String()
	case lang.Italian:
		result = it.Effectiveness(e).String()
	case lang.Korean:
		result = ko.Effectiveness(e).String()
	case lang.Spanish:
		result = es.Effectiveness(e).String()
	case lang.ChineseSimplified:
		result = zhCHS.Effectiveness(e).String()
	case lang.ChineseTraditional:
		result = zhCHT.Effectiveness(e).String()
	case lang.ChineseChina:
		result = zhCN.Effectiveness(e).String()
	case lang.ChineseHongKong:
		result = zhHK.Effectiveness(e).String()
	case lang.ChineseTaiwan:
		result = zhTW.Effectiveness(e).String()
	}
	result = strings.Replace(result, "_", " ", -1)
	return result
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
