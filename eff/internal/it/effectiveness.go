package it

// Effectiveness
type Effectiveness int8

const (
	Nessun_Effetto Effectiveness = iota
	Non_Molto_Efficace
	Normale
	Super_Efficace
)

//go:generate stringer -type=Effectiveness
