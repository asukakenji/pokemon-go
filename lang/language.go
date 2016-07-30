package lang

// Language
type Language int8

const (
	Japanese Language = iota
	English
	French
	German
	Italian
	Korean
	Spanish
	ChineseSimplified
	ChineseTraditional
	ChineseChina
	ChineseHongKong
	ChineseTaiwan
)

//go:generate stringer -type=Language

func (l Language) LocalName() string {
	switch l {
	case Japanese:
		return "日本語"
	case English:
		return "English"
	case French:
		return "Français"
	case German:
		return "Deutsch"
	case Italian:
		return "Italiano"
	case Korean:
		return "한국어"
	case Spanish:
		return "Español"
	case ChineseSimplified:
		return "简体中文（官方译名）"
	case ChineseTraditional:
		return "繁體中文（官方譯名）"
	case ChineseChina:
		return "简体中文（中国译名）"
	case ChineseHongKong:
		return "繁體中文（香港譯名）"
	case ChineseTaiwan:
		return "繁體中文（台灣譯名）"
	default:
		return ""
	}
}

// Reference:
// http://publications.europa.eu/mdr/resource/authority/multilingual/xml/multilingual.xml

func (l Language) ShortName() string {
	switch l {
	case Japanese:
		return "日"
	case English:
		return "Eng"
	case French:
		return "Fra"
	case German:
		return "Deu"
	case Italian:
		return "Ita"
	case Korean:
		return "한"
	case Spanish:
		return "Spa"
	case ChineseSimplified:
		return "简(官)"
	case ChineseTraditional:
		return "繁(官)"
	case ChineseChina:
		return "简(中)"
	case ChineseHongKong:
		return "繁(港)"
	case ChineseTaiwan:
		return "繁(台)"
	default:
		return ""
	}
}
