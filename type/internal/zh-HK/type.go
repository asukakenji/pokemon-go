package zhHK

// Type
type Type int8

const (
	沒有  Type = iota // -
	普通              // OK: https://youtu.be/ufS3IjJwXY0?t=43s
	火               // OK: https://youtu.be/dFQ3cr6lY7Q?t=3m3s , https://youtu.be/4Ygday15XxM?t=27s
	水               // OK: https://youtu.be/EkCoEz8FebU?t=4m53s
	草               //
	電擊              // OK: https://youtu.be/ufS3IjJwXY0?t=46s , https://youtu.be/aKhuwCZuWgg?t=4m13s
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
	鋼               // OK: https://youtu.be/dFQ3cr6lY7Q?t=3m1s , https://youtu.be/4Ygday15XxM?t=25s , https://youtu.be/ufS3IjJwXY0?t=1m57s
	妖精              //
)

//go:generate stringer -type=Type
