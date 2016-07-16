package fr

// Effectiveness
type Effectiveness int8

const (
	Aucun_Effet Effectiveness = iota
	Pas_Très_Efficace
	Normaux
	Super_Efficace
)

//go:generate stringer -type=Effectiveness
