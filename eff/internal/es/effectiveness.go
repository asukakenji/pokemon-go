package es

// Effectiveness
type Effectiveness int8

const (
	Sin_Efecto Effectiveness = iota
	No_Muy_Eficaz
	Normal
	Súper_Eficaz
)

//go:generate stringer -type=Effectiveness
