package pokemon

import (
	"github.com/asukakenji/pokemon-go/lang"

	"github.com/asukakenji/pokemon-go/internal/type/de"
	"github.com/asukakenji/pokemon-go/internal/type/en"
	"github.com/asukakenji/pokemon-go/internal/type/es"
	"github.com/asukakenji/pokemon-go/internal/type/fr"
	"github.com/asukakenji/pokemon-go/internal/type/it"
	"github.com/asukakenji/pokemon-go/internal/type/ja"
	"github.com/asukakenji/pokemon-go/internal/type/ko"
	"github.com/asukakenji/pokemon-go/internal/type/zh-CN"
	"github.com/asukakenji/pokemon-go/internal/type/zh-HK"
	"github.com/asukakenji/pokemon-go/internal/type/zh-TW"
)

// Type
type Type int8

const (
	TypeNone     Type = iota // -
	TypeNormal               // 無
	TypeFire                 // 炎
	TypeWater                // 水
	TypeGrass                // 草
	TypeElectric             // 電
	TypeIce                  // 氷
	TypeFighting             // 闘
	TypePoison               // 毒
	TypeGround               // 地
	TypeFlying               // 飛
	TypePsychic              // 超
	TypeBug                  // 虫
	TypeRock                 // 岩
	TypeGhost                // 霊
	TypeDragon               // 竜
	TypeDark                 // 悪
	TypeSteel                // 鋼
	TypeFairy                // 妖
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

const _Type_Symbol_string = "無炎水草電氷闘毒地飛超虫岩霊竜悪鋼妖"

var _Type_Symbol_index = [...]uint8{0, 0, 3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54}

func (t Type) Symbol() string {
	if t < 0 || t >= Type(len(_Type_Symbol_index)-1) {
		panic("Invalid Type")
	}
	return _Type_Symbol_string[_Type_Symbol_index[t]:_Type_Symbol_index[t+1]]
}
