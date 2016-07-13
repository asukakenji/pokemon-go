package es

// Type
type Type int8

const (
	Ninguna Type = iota
	Normal
	Fuego
	Agua
	Planta // Hierba
	Eléctrico
	Hielo
	Lucha // Peleador
	Veneno
	Tierra
	Volador
	Psíquico
	Insecto // Bicho
	Roca
	Fantasma
	Dragón
	Siniestro // Oscuro
	Acero
	Hada
)

//go:generate stringer -type=Type
