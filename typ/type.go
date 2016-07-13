package typ

import (
	"github.com/asukakenji/pokemon-go/lang"

	"github.com/asukakenji/pokemon-go/typ/internal/de"
	"github.com/asukakenji/pokemon-go/typ/internal/en"
	"github.com/asukakenji/pokemon-go/typ/internal/es"
	"github.com/asukakenji/pokemon-go/typ/internal/fr"
	"github.com/asukakenji/pokemon-go/typ/internal/it"
	"github.com/asukakenji/pokemon-go/typ/internal/ja"
	"github.com/asukakenji/pokemon-go/typ/internal/ko"
	"github.com/asukakenji/pokemon-go/typ/internal/zh-CN"
	"github.com/asukakenji/pokemon-go/typ/internal/zh-HK"
	"github.com/asukakenji/pokemon-go/typ/internal/zh-TW"
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

func (t Type) Id() int {
	return int(t)
}

func (t Type) LocalName(l lang.Language) string {
	switch l {
	case lang.Japanese:
		return ja.Type(t).String()
	case lang.English:
		return en.Type(t).String()
	case lang.French:
		return fr.Type(t).String()
	case lang.German:
		return de.Type(t).String()
	case lang.Italian:
		return it.Type(t).String()
	case lang.Korean:
		return ko.Type(t).String()
	case lang.Spanish:
		return es.Type(t).String()
	case lang.ChineseChina:
		return zhCN.Type(t).String()
	case lang.ChineseHongKong:
		return zhHK.Type(t).String()
	case lang.ChineseTaiwan:
		return zhTW.Type(t).String()
	default:
		return ""
	}
}

const _Symbol_string = "無炎水草電氷闘毒地飛超虫岩霊竜悪鋼妖"

var _Symbol_index = [...]uint8{0, 0, 3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54}

func (t Type) Symbol() string {
	if t < 0 || t >= Type(len(_Symbol_index)-1) {
		panic("Invalid Type")
	}
	return _Symbol_string[_Symbol_index[t]:_Symbol_index[t+1]]
}
