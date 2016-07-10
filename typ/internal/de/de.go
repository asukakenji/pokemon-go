package de

// Type
type Type int8

const (
	Normal Type = iota
	Feuer
	Wasser
	Pflanze
	Elektro
	Eis
	Kampf
	Gift
	Boden
	Flug
	Psycho
	Käfer
	Gestein
	Geist
	Drache
	Unlicht
	Stahl
	Fee
)

//go:generate stringer -type=Type
