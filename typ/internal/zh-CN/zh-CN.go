package zhCN

// Type
type Type int8

const (
	沒有 Type = iota
	一般
	炎
	水
	草
	电
	冰
	格斗
	毒
	地上
	飞行
	超能力
	虫
	岩石
	幽灵
	龙
	恶
	钢铁
	妖精
)

//go:generate stringer -type=Type
