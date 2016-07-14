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
	ChineseChina
	ChineseHongKong
	// ChineseMacau
	// ChineseSingapore
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
	case ChineseChina:
		return "中文（中国）"
	case ChineseHongKong:
		return "中文（香港）"
	case ChineseTaiwan:
		return "中文（台灣）"
	default:
		return ""
	}
}
