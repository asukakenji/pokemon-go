package fr

// Type
type Type int8

const (
	Aucun Type = iota
	Normal
	Feu
	Eau
	Plante
	Électrique
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
	Ténèbres
	Acier
	Fée
)

//go:generate stringer -type=Type
