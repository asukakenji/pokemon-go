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

// The "T" prefix is required to make the symbols public.
// Go requires the first character of a public identifier to be in uppercase.
const (
	T無 Type = iota // Normal
	T炎             // Fire
	T水             // Water
	T草             // Grass
	T電             // Electric
	T氷             // Ice
	T闘             // Fighting
	T毒             // Poison
	T地             // Ground
	T飛             // Flying
	T超             // Psychic
	T虫             // Bug
	T岩             // Rock
	T霊             // Ghost
	T竜             // Dragon
	T悪             // Dark
	T鋼             // Steel
	T妖             // Fairy
)

//go:generate stringer -type=Type

func ForEach(f func(Type)) {
	for i := T無; i <= T妖; i = i + 1 {
		f(Type(i))
	}
}

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

var _Symbol_index = [...]uint8{0, 3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54}

func (t Type) Symbol() string {
	if t < 0 || t >= Type(len(_Symbol_index)-1) {
		panic("Invalid Type")
	}
	return _Symbol_string[_Symbol_index[t]:_Symbol_index[t+1]]
}
