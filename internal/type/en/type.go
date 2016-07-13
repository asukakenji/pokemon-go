package en

// Type
type Type int8

const (
	None Type = iota
	Normal
	Fire
	Water
	Grass
	Electric
	Ice
	Fighting
	Poison
	Ground
	Flying
	Psychic
	Bug
	Rock
	Ghost
	Dragon
	Dark
	Steel
	Fairy
)

//go:generate stringer -type=Type
