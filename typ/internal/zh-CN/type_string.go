// Code generated by "stringer -type=Type"; DO NOT EDIT

package zhCN

import "fmt"

const _Type_name = "一般炎水草电冰格斗毒地上飞行超能力虫岩石幽灵龙恶钢铁妖精"

var _Type_index = [...]uint8{0, 6, 9, 12, 15, 18, 21, 27, 30, 36, 42, 51, 54, 60, 66, 69, 72, 78, 84}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
