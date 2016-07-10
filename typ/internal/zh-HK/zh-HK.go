package zhHK

// Type
type Type int8

const (
	普通  Type = iota // OK
	火               //
	水               // OK: https://youtu.be/EkCoEz8FebU?t=4m53s
	草               //
	電擊              // OK
	冰               //
	格鬥              //
	毒               //
	地上              //
	飛行              //
	超能力             //
	蟲               //
	岩石              //
	幽靈              //
	龍               //
	惡               //
	鋼               //
	妖精              //
)

//go:generate stringer -type=Type
