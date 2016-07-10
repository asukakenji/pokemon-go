package fr

// Type
type Type int8

const (
	Normal Type = iota
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
