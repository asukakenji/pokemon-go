package _type

import (
	"strings"

	"github.com/asukakenji/pokemon-go/eff"
	"github.com/asukakenji/pokemon-go/lang"

	"github.com/asukakenji/pokemon-go/type/internal/de"
	"github.com/asukakenji/pokemon-go/type/internal/en"
	"github.com/asukakenji/pokemon-go/type/internal/es"
	"github.com/asukakenji/pokemon-go/type/internal/fr"
	"github.com/asukakenji/pokemon-go/type/internal/it"
	"github.com/asukakenji/pokemon-go/type/internal/ja"
	"github.com/asukakenji/pokemon-go/type/internal/ko"
	"github.com/asukakenji/pokemon-go/type/internal/zh-CHS"
	"github.com/asukakenji/pokemon-go/type/internal/zh-CHT"
	"github.com/asukakenji/pokemon-go/type/internal/zh-CN"
	"github.com/asukakenji/pokemon-go/type/internal/zh-HK"
	"github.com/asukakenji/pokemon-go/type/internal/zh-TW"
)

// Type
type Type int8

const (
	None     Type = iota // -
	Normal               // 無
	Fire                 // 炎
	Water                // 水
	Grass                // 草
	Electric             // 電
	Ice                  // 氷
	Fighting             // 闘
	Poison               // 毒
	Ground               // 地
	Flying               // 飛
	Psychic              // 超
	Bug                  // 虫
	Rock                 // 岩
	Ghost                // 霊
	Dragon               // 竜
	Dark                 // 悪
	Steel                // 鋼
	Fairy                // 妖
)

//go:generate stringer -type=Type

func PackageName(l lang.Language) string {
	switch l {
	case lang.Japanese:
		return "タイプ"
	case lang.English:
		return "Type"
	case lang.French:
		return "Type"
	case lang.German:
		return "Typ"
	case lang.Italian:
		return "Tipo"
	case lang.Korean:
		return "타입"
	case lang.Spanish:
		return "Tipo"
	case lang.ChineseSimplified:
		return "属性"
	case lang.ChineseTraditional:
		return "屬性"
	case lang.ChineseChina:
		return "属性"
	case lang.ChineseHongKong:
		return "屬性"
	case lang.ChineseTaiwan:
		return "屬性"
	default:
		return ""
	}
}

func (t Type) Id() int {
	return int(t)
}

func (t Type) IsValid() bool {
	return Normal <= t && t <= Fairy
}

func (t Type) LocalName(l lang.Language) string {
	result := ""
	switch l {
	case lang.Japanese:
		result = ja.Type(t).String()
	case lang.English:
		result = en.Type(t).String()
	case lang.French:
		result = fr.Type(t).String()
	case lang.German:
		result = de.Type(t).String()
	case lang.Italian:
		result = it.Type(t).String()
	case lang.Korean:
		result = ko.Type(t).String()
	case lang.Spanish:
		result = es.Type(t).String()
	case lang.ChineseSimplified:
		result = zhCHS.Type(t).String()
	case lang.ChineseTraditional:
		result = zhCHT.Type(t).String()
	case lang.ChineseChina:
		result = zhCN.Type(t).String()
	case lang.ChineseHongKong:
		result = zhHK.Type(t).String()
	case lang.ChineseTaiwan:
		result = zhTW.Type(t).String()
	}
	result = strings.Replace(result, "_full_stop_", ".", -1)
	return result
}

// t: The type of the pokemon (defender).
// moveType: The type of the move (attacker);
func (t Type) AttackedBy(moveType Type) eff.Effectiveness {
	return effectivenesses[t][moveType]
}

const _Symbol_string = "無炎水草電氷闘毒地飛超虫岩霊竜悪鋼妖"

var _Symbol_index = [...]uint8{0, 0, 3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54}

func (t Type) Symbol() string {
	if t < 0 || t >= Type(len(_Symbol_index)-1) {
		panic("Invalid Type")
	}
	return _Symbol_string[_Symbol_index[t]:_Symbol_index[t+1]]
}
