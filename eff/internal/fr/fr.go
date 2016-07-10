package fr

// Effectiveness
type Effectiveness int8

const (
	無効 Effectiveness = iota
	今一つ
	通常
	抜群
)

//go:generate stringer -type=Effectiveness
