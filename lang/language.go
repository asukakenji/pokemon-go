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

func ForEach(f func(Language)) {
	for i := Japanese; i <= ChineseTaiwan; i = i + 1 {
		f(Language(i))
	}
}
