// Code generated by "stringer -type=Language"; DO NOT EDIT

package lang

import "fmt"

const _Language_name = "JapaneseEnglishFrenchGermanItalianKoreanSpanishChineseSimplifiedChineseTraditionalChineseChinaChineseHongKongChineseTaiwan"

var _Language_index = [...]uint8{0, 8, 15, 21, 27, 34, 40, 47, 64, 82, 94, 109, 122}

func (i Language) String() string {
	i -= 1
	if i < 0 || i >= Language(len(_Language_index)-1) {
		return fmt.Sprintf("Language(%d)", i+1)
	}
	return _Language_name[_Language_index[i]:_Language_index[i+1]]
}
