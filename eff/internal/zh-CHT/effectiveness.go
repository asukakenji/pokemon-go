package zhCHT

// Effectiveness
type Effectiveness int8

const (
	無效 Effectiveness = iota
	不佳
	普通
	超群
)

//go:generate stringer -type=Effectiveness
