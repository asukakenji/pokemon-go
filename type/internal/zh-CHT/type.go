package zhCHT

// Type
type Type int8

const (
	沒有 Type = iota
	一般
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
