package it

// Type
type Type int8

const (
	Normale Type = iota
	Fuoco
	Acqua
	Erba
	Elettro
	Ghiaccio
	Lotta
	Veleno
	Terra
	Volante
	Psico
	Coleottero
	Roccia
	Spettro
	Drago
	Buio
	Acciaio
	Folletto
)

//go:generate stringer -type=Type
