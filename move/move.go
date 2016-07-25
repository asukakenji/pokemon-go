package move

import (
	"fmt"

	typ "github.com/asukakenji/pokemon-go/type"
)

// Move
type Move interface {
	fmt.Stringer
	Id() int
	Type() typ.Type
	Damage() int
	EnergyRequired() int
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

func (m StandardMove) Id() int {
	return int(m)
}

func (m StandardMove) Type() typ.Type {
	return m.self()._type
}

func (m StandardMove) Damage() int {
	return int(m.self().damage)
}

func (m StandardMove) EnergyRequired() int {
	return 0
}

func (m StandardMove) self() *_standard_move {
	return moves[m.Id()].(*_standard_move)
}

func (m SpecialMove) Id() int {
	return int(m)
}

func (m SpecialMove) Type() typ.Type {
	return m.self()._type
}

func (m SpecialMove) Damage() int {
	return int(m.self().damage)
}

func (m SpecialMove) EnergyRequired() int {
	return int(m.self().energyRequired)
}

func (m SpecialMove) self() *_special_move {
	return moves[m.Id()].(*_special_move)
}
