package de

// Effectiveness
type Effectiveness int8

const (
	Kein_Effekt Effectiveness = iota
	Nicht_Sehr_Effektiv
	Normaler
	Sehr_Effektiv
)

//go:generate stringer -type=Effectiveness
