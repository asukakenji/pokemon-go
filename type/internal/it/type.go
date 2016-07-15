package it

// Type
type Type int8

const (
	Nessuna Type = iota
	Normale
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
	Coleott_full_stop_
	Roccia
	Spettro
	Drago
	Buio
	Acciaio
	Folletto
)

//go:generate stringer -type=Type
