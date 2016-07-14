package fr

// Type
type Type int8

const (
	Aucun Type = iota
	Normal
	Feu
	Eau
	Plante
	Électrik
	Glace
	Combat
	Poison
	Sol
	Vol
	Psy
	Insecte
	Roche
	Spectre
	Dragon
	Ténèbr_full_stop_
	Acier
	Fée
)

//go:generate stringer -type=Type
