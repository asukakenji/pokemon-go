package move

import (
	"fmt"

	typ "github.com/asukakenji/pokemon-go/type"
)

// Move
type Move interface {
	fmt.Stringer

	// Id returns the id of this move.
	Id() int

	// IsValid returns whether this move is valid.
	// All moves declared in this package are valid.
	// If this move is casted from an integer, it may be invalid.
	IsValid() bool

	// Type returns the type of this move.
	Type() typ.Type

	// Damage returns the energy required by this move.
	Damage() int

	// EnergyIncreased returns the energy increased after making this move.
	EnergyIncreased() int

	// EnergyRequired returns the energy required to make this move.
	EnergyRequired() int

	// CriticalHitChance returns the chance of critical hit in percent.
	CriticalHitChance() int

	// Duration returns the duration of this move in seconds.
	Duration() float64
}

type StandardMove int16
type SpecialMove int16

const (
	_ StandardMove = iota
	Acid
	Bite
	Bubble
	BugBite
	BulletPunch
	Confusion
	Cut
	DragonBreath
	Ember
	FeintAttack
	FireFang
	FrostBreath
	FuryCutter
	IceShard
	KarateChop
	Lick
	LowKick
	MetalClaw
	MudShot
	MudSlap
	Peck
	PoisonJab
	PoisonSting
	Pound
	PsychoCut
	QuickAttack
	RazorLeaf
	RockSmash
	RockThrow
	Scratch
	ShadowClaw
	Spark
	Splash
	SteelWing
	SuckerPunch
	Tackle
	ThunderShock
	VineWhip
	WaterGun
	WingAttack
	ZenHeadbutt
	AerialAce SpecialMove = iota
	AirCutter
	AncientPower
	AquaJet
	AquaTail
	Blizzard
	BodySlam
	BoneClub
	BrickBreak
	Brine
	BubbleBeam
	BugBuzz
	Bulldoze
	CrossChop
	CrossPoison
	DarkPulse
	DazzlingGleam
	Dig
	DisarmingVoice
	Discharge
	DragonClaw
	DragonPulse
	DrainingKiss
	DrillPeck
	DrillRun
	Earthquake
	FireBlast
	FirePunch
	FlameBurst
	FlameCharge
	FlameWheel
	Flamethrower
	FlashCannon
	GunkShot
	HeatWave
	HornAttack
	Hurricane
	HydroPump
	HyperBeam
	HyperFang
	IceBeam
	IcePunch
	IcyWind
	IronHead
	LeafBlade
	LowSweep
	MagnetBomb
	Megahorn
	Moonblast
	MudBomb
	NightSlash
	OminousWind
	PetalBlizzard
	PlayRough
	PoisonFang
	PowerGem
	PowerWhip
	Psybeam
	Psychic
	Psyshock
	RockSlide
	RockTomb
	Scald
	SeedBomb
	ShadowBall
	SignalBeam
	Sludge
	SludgeBomb
	SludgeWave
	SolarBeam
	Stomp
	StoneEdge
	Struggle
	Submission
	Swift
	Thunder
	ThunderPunch
	Thunderbolt
	Twister
	ViceGrip
	WaterPulse
	Wrap
	X_Scissor
)

//go:generate stringer -type=StandardMove
//go:generate stringer -type=SpecialMove

// Id returns the id of this move.
func (m StandardMove) Id() int {
	return int(m)
}

// IsValid returns whether this move is valid.
// All moves declared in this package are valid.
// If this move is casted from an integer, it may be invalid.
func (m StandardMove) IsValid() bool {
	return Acid <= m && m <= ZenHeadbutt
}

// Type returns the type of this move.
func (m StandardMove) Type() typ.Type {
	return m.self()._type
}

// Damage returns the damage of this move.
func (m StandardMove) Damage() int {
	return int(m.self().damage)
}

// EnergyIncreased returns the energy increased after making this move.
func (m StandardMove) EnergyIncreased() int {
	return int(m.self().energyIncreased)
}

// EnergyRequired returns the energy required to make this move.
// All standard moves return 0.
func (m StandardMove) EnergyRequired() int {
	return 0
}

// CriticalHitChance returns the chance of critical hit in percent.
func (m StandardMove) CriticalHitChance() int {
	return 0
}

// Duration returns the duration of this move in seconds.
func (m StandardMove) Duration() float64 {
	return float64(m.self().duration)
}

// self returns the actual move object.
func (m StandardMove) self() *_standard_move {
	return moves[m.Id()].(*_standard_move)
}

// Id returns the id of this move.
func (m SpecialMove) Id() int {
	return int(m)
}

// IsValid returns whether this move is valid.
// All moves declared in this package are valid.
// If this move is casted from an integer, it may be invalid.
func (m SpecialMove) IsValid() bool {
	return AerialAce <= m && m <= X_Scissor
}

// Type returns the type of this move.
func (m SpecialMove) Type() typ.Type {
	return m.self()._type
}

// Damage returns the energy required by this move.
func (m SpecialMove) Damage() int {
	return int(m.self().damage)
}

// EnergyIncreased returns the energy increased after making this move.
// All special moves return 0.
func (m SpecialMove) EnergyIncreased() int {
	return 0
}

// EnergyRequired returns the energy required to make this move.
func (m SpecialMove) EnergyRequired() int {
	return int(100.0 / m.self().moveCountPerFullEnergy)
}

// CriticalHitChance returns the chance of critical hit in percent.
func (m SpecialMove) CriticalHitChance() int {
	return int(m.self().criticalHitChance)
}

// Duration returns the duration of this move in seconds.
func (m SpecialMove) Duration() float64 {
	return float64(m.self().duration)
}

// self returns the actual move object.
func (m SpecialMove) self() *_special_move {
	return moves[m.Id()].(*_special_move)
}
