package lv

import (
	"fmt"
	"math"
)

// Level
type Level float32

func (l Level) Id() float32 {
	return float32(l)
}

func (l Level) String() string {
	return fmt.Sprintf("%.1f", l.Id())
}

func (l Level) IsValid() bool {
	return 1.0 <= l && l <= 40.0 && (l.IsInteger() || (l * 2).IsInteger())
}

func (l Level) IsInteger() bool {
	f := float64(l)
	return !math.IsInf(f, 0) && !math.IsNaN(f) && l == Level(float32(math.Trunc(f)))
}

func (l Level) StardustToPowerUp() int {
	return int(l.self().stardustToPowerUp)
}

func (l Level) CandyToPowerUp() int {
	return int(l.self().candyToPowerUp)
}

func (l Level) StardustAndCandy() (int, int) {
	return l.StardustToPowerUp(), l.CandyToPowerUp()
}

func (l Level) CombatPowerMultiplier() float64 {
	return l.self().combatPowerMultiplier
}

func (l Level) self() *_level {
	return levels[int((l-1.0)*2)]
}

func LevelsByStardust(stardust int) []Level {
	result := make([]Level, 0)
	for _, level := range levels {
		if int(level.stardustToPowerUp) == stardust {
			result = append(result, level.id)
		}
	}
	return result
}

func LevelsByStardustAndCandy(stardust, candy int) []Level {
	result := make([]Level, 0)
	for _, level := range levels {
		if int(level.stardustToPowerUp) == stardust && int(level.candyToPowerUp) == candy {
			result = append(result, level.id)
		}
	}
	return result
}
