package en

// Effectiveness
type Effectiveness int8

const (
	No_Effect Effectiveness = iota
	Not_Very_Effective
	Normal
	Super_Effective
)

//go:generate stringer -type=Effectiveness
