// Code generated by "stringer -type=Type"; DO NOT EDIT

package de

import "fmt"

const _Type_name = "KeinerNormalFeuerWasserPflanzeElektroEisKampfGiftBodenFlugPsychoKäferGesteinGeistDracheUnlichtStahlFee"

var _Type_index = [...]uint8{0, 6, 12, 17, 23, 30, 37, 40, 45, 49, 54, 58, 64, 70, 77, 82, 88, 95, 100, 103}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
