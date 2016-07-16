package ko

// Effectiveness
type Effectiveness int8

// Reference:
// https://github.com/contolini/pokedex/blob/master/data/csv/ability_flavor_text.csv
const (
	효과가없다 Effectiveness = iota // Google Translate
	별로인                        // REF#110 , 별로인 , 효과가별로인
	표준                         // Google Translate
	굉장한                        // REF#25 , 굉장한 , 효과가굉장한
)

//go:generate stringer -type=Effectiveness
