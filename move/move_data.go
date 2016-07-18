package move

import (
	t "github.com/asukakenji/pokemon-go/type"
)

type _standard_move struct {
	id     StandardMove
	_type  t.Type
	damage int16
}

type _special_move struct {
	id             SpecialMove
	_type          t.Type
	damage         int16
	energyRequired int8
}

var moves = [...]interface{}{
	nil,
	// Standard Move
	&_standard_move{Acid, t.Poison, 10},
	&_standard_move{Bite, t.Dark, 6},
	&_standard_move{Bubble, t.Water, 15},
	&_standard_move{BugBite, t.Bug, 6},
	&_standard_move{BulletPunch, t.Steel, 10},
	&_standard_move{Confusion, t.Psychic, 12},
	&_standard_move{DragonBreath, t.Dragon, 6},
	&_standard_move{Ember, t.Fire, 10},
	&_standard_move{FeintAttack, t.Dark, 12},
	&_standard_move{FireFang, t.Fire, 7},
	&_standard_move{FrostBreath, t.Ice, 12},
	&_standard_move{FuryCutter, t.Bug, 3},
	&_standard_move{IceShard, t.Ice, 10},
	&_standard_move{KarateChop, t.Fighting, 6},
	&_standard_move{Lick, t.Ghost, 10},
	&_standard_move{LowKick, t.Fighting, 5},
	&_standard_move{MetalClaw, t.Steel, 12},
	&_standard_move{MudShot, t.Ground, 12},
	&_standard_move{MudSlap, t.Ground, 6},
	&_standard_move{Peck, t.Flying, 10},
	&_standard_move{PoisonJab, t.Poison, 15},
	&_standard_move{PoisonSting, t.Poison, 6},
	&_standard_move{Pound, t.Normal, 8},
	&_standard_move{PsychoCut, t.Psychic, 15},
	&_standard_move{QuickAttack, t.Normal, 10},
	&_standard_move{RazorLeaf, t.Grass, 15},
	&_standard_move{RockSmash, t.Fighting, 5},
	&_standard_move{RockThrow, t.Rock, 12},
	&_standard_move{Scratch, t.Normal, 10},
	&_standard_move{ShadowClaw, t.Ghost, 16},
	&_standard_move{Spark, t.Electric, 7},
	&_standard_move{Splash, t.Water, 0},
	&_standard_move{SteelWing, t.Steel, 15},
	&_standard_move{SuckerPunch, t.Dark, 7},
	&_standard_move{Tackle, t.Normal, 12},
	&_standard_move{ThunderShock, t.Electric, 5},
	&_standard_move{VineWhip, t.Grass, 10},
	&_standard_move{WaterGun, t.Water, 10},
	&_standard_move{WingAttack, t.Flying, 12},
	&_standard_move{ZenHeadbutt, t.Psychic, 15},
	// Special Move
	&_special_move{AerialAce, t.Flying, 25, 4},
	&_special_move{AirCutter, t.Flying, 25, 4},
	&_special_move{AncientPower, t.Rock, 30, 4},
	&_special_move{AquaJet, t.Water, 15, 5},
	&_special_move{AquaTail, t.Water, 50, 2},
	&_special_move{Blizzard, t.Ice, 60, 1},
	&_special_move{BodySlam, t.Normal, 50, 2},
	&_special_move{BoneClub, t.Ground, 20, 4},
	&_special_move{BrickBreak, t.Fighting, 30, 3},
	&_special_move{Brine, t.Water, 15, 5},
	&_special_move{BubbleBeam, t.Water, 25, 4},
	&_special_move{BugBuzz, t.Bug, 50, 2},
	&_special_move{Bulldoze, t.Ground, 30, 4},
	&_special_move{CrossChop, t.Fighting, 55, 2},
	&_special_move{CrossPoison, t.Poison, 20, 4},
	&_special_move{DarkPulse, t.Dark, 45, 3},
	&_special_move{DazzlingGleam, t.Fairy, 45, 3},
	&_special_move{Dig, t.Ground, 45, 3},
	&_special_move{DisarmingVoice, t.Fairy, 20, 5},
	&_special_move{Discharge, t.Electric, 40, 3},
	&_special_move{DragonClaw, t.Dragon, 40, 2},
	&_special_move{DragonPulse, t.Dragon, 50, 2},
	&_special_move{DrainingKiss, t.Fairy, 15, 5},
	&_special_move{DrillPeck, t.Flying, 30, 3},
	&_special_move{DrillRun, t.Ground, 40, 3},
	&_special_move{Earthquake, t.Ground, 60, 1},
	&_special_move{FireBlast, t.Fire, 60, 1},
	&_special_move{FirePunch, t.Fire, 35, 3},
	&_special_move{FlameBurst, t.Fire, 25, 4},
	&_special_move{FlameCharge, t.Fire, 25, 4},
	&_special_move{FlameWheel, t.Fire, 35, 4},
	&_special_move{Flamethrower, t.Fire, 50, 2},
	&_special_move{FlashCannon, t.Steel, 55, 3},
	&_special_move{GunkShot, t.Poison, 60, 1},
	&_special_move{HeatWave, t.Fire, 60, 1},
	&_special_move{HornAttack, t.Normal, 20, 4},
	&_special_move{Hurricane, t.Flying, 60, 1},
	&_special_move{HydroPump, t.Water, 60, 1},
	&_special_move{HyperBeam, t.Normal, 70, 1},
	&_special_move{HyperFang, t.Normal, 35, 3},
	&_special_move{IceBeam, t.Ice, 50, 2},
	&_special_move{IcePunch, t.Ice, 45, 3},
	&_special_move{IcyWind, t.Ice, 15, 5},
	&_special_move{IronHead, t.Steel, 40, 3},
	&_special_move{LeafBlade, t.Grass, 45, 2},
	&_special_move{LowSweep, t.Fighting, 25, 4},
	&_special_move{MagnetBomb, t.Steel, 25, 4},
	&_special_move{Megahorn, t.Bug, 55, 1},
	&_special_move{Moonblast, t.Fairy, 60, 1},
	&_special_move{MudBomb, t.Ground, 25, 4},
	&_special_move{NightSlash, t.Dark, 25, 4},
	&_special_move{OminousWind, t.Ghost, 25, 4},
	&_special_move{PetalBlizzard, t.Grass, 50, 2},
	&_special_move{PlayRough, t.Fairy, 50, 2},
	&_special_move{PoisonFang, t.Poison, 15, 5},
	&_special_move{PowerGem, t.Rock, 40, 3},
	&_special_move{PowerWhip, t.Grass, 60, 1},
	&_special_move{Psybeam, t.Psychic, 35, 4},
	&_special_move{Psychic, t.Psychic, 50, 2},
	&_special_move{Psyshock, t.Psychic, 40, 3},
	&_special_move{RockSlide, t.Rock, 40, 3},
	&_special_move{RockTomb, t.Rock, 25, 4},
	&_special_move{Scald, t.Water, 35, 3},
	&_special_move{SeedBomb, t.Grass, 30, 3},
	&_special_move{ShadowBall, t.Ghost, 40, 3},
	&_special_move{SignalBeam, t.Bug, 35, 3},
	&_special_move{Sludge, t.Poison, 25, 4},
	&_special_move{SludgeBomb, t.Poison, 50, 2},
	&_special_move{SludgeWave, t.Poison, 60, 1},
	&_special_move{SolarBeam, t.Grass, 65, 1},
	&_special_move{Stomp, t.Normal, 25, 4},
	&_special_move{StoneEdge, t.Rock, 55, 1},
	&_special_move{Struggle, t.Normal, 15, 5},
	&_special_move{Submission, t.Fighting, 30, 3},
	&_special_move{Swift, t.Normal, 25, 4},
	&_special_move{Thunder, t.Electric, 65, 1},
	&_special_move{ThunderPunch, t.Electric, 40, 3},
	&_special_move{Thunderbolt, t.Electric, 50, 2},
	&_special_move{Twister, t.Dragon, 15, 5},
	&_special_move{ViceGrip, t.Normal, 15, 5},
	&_special_move{WaterPulse, t.Water, 30, 4},
	&_special_move{Wrap, t.Normal, 15, 5},
	&_special_move{X_Scissor, t.Bug, 30, 3},
}
