package zhTW

// Type
type Type int8

const (
	一般 Type = iota
	炎
	水
	草
	電
	冰
	格鬥
	毒
	地上
	飛行
	超能力
	蟲
	岩石
	幽靈
	龍
	惡
	鋼鐵
	妖精
)

//go:generate stringer -type=Type
