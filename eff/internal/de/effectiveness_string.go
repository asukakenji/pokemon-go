// Code generated by "stringer -type=Effectiveness"; DO NOT EDIT

package de

import "fmt"

const _Effectiveness_name = "無効今一つ通常抜群"

var _Effectiveness_index = [...]uint8{0, 6, 15, 21, 27}

func (i Effectiveness) String() string {
	if i < 0 || i >= Effectiveness(len(_Effectiveness_index)-1) {
		return fmt.Sprintf("Effectiveness(%d)", i)
	}
	return _Effectiveness_name[_Effectiveness_index[i]:_Effectiveness_index[i+1]]
}
