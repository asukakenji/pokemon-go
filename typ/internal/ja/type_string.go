// Code generated by "stringer -type=Type"; DO NOT EDIT

package ja

import "fmt"

const _Type_name = "ノーマルほのおみずくさでんきこおりかくとうどくじめんひこうエスパーむしいわゴーストドラゴンあくはがねフェアリー"

var _Type_index = [...]uint8{0, 12, 21, 27, 33, 42, 51, 63, 69, 78, 87, 99, 105, 111, 123, 135, 141, 150, 165}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
