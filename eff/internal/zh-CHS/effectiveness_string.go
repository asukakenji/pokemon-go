// Code generated by "stringer -type=Effectiveness"; DO NOT EDIT

package zhCHS

import "fmt"

const _Effectiveness_name = "无效不佳普通超群"

var _Effectiveness_index = [...]uint8{0, 6, 12, 18, 24}

func (i Effectiveness) String() string {
	if i < 0 || i >= Effectiveness(len(_Effectiveness_index)-1) {
		return fmt.Sprintf("Effectiveness(%d)", i)
	}
	return _Effectiveness_name[_Effectiveness_index[i]:_Effectiveness_index[i+1]]
}
