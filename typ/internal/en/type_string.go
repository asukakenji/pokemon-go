// Code generated by "stringer -type=Type"; DO NOT EDIT

package en

import "fmt"

const _Type_name = "NoneNormalFireWaterGrassElectricIceFightingPoisonGroundFlyingPsychicBugRockGhostDragonDarkSteelFairy"

var _Type_index = [...]uint8{0, 4, 10, 14, 19, 24, 32, 35, 43, 49, 55, 61, 68, 71, 75, 80, 86, 90, 95, 100}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
