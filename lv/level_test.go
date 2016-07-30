package lv

import (
	"math"
	"testing"
)

// TestIsValid checks the method IsValid() with predefined test cases.
func TestIsValid(t *testing.T) {
	cases := []struct {
		level Level
		expected bool
	}{
		{Level(-2.0), false},
		{Level(-1.5), false},
		{Level(-1.0), false},
		{Level(-0.5), false},
		{Level(0.0), false},
		{Level(0.5), false},
		{Level(1.0), true},
		{Level(1.25), false},
		{Level(1.5), true},
		{Level(20.0), true},
		{Level(20.5), true},
		{Level(39.5), true},
		{Level(40.0), true},
		{Level(40.5), false},
		{Level(math.Inf(-1)), false},
		{Level(math.Inf(1)), false},
		{Level(math.NaN()), false},
	}
	for _, c := range cases {
		got := c.level.IsValid()
		if got != c.expected {
			t.Errorf("Level(%f).IsValid() = %t", c.level, got)
		}
	}
}

// TestIsInteger checks the method IsInteger() with predefined test cases.
func TestIsInteger(t *testing.T) {
	cases := []struct {
		level Level
		expected bool
	}{
		{Level(-2.0), true},
		{Level(-1.5), false},
		{Level(-1.0), true},
		{Level(-0.5), false},
		{Level(0.0), true},
		{Level(0.5), false},
		{Level(1.0), true},
		{Level(1.25), false},
		{Level(1.5), false},
		{Level(20.0), true},
		{Level(20.5), false},
		{Level(39.5), false},
		{Level(40.0), true},
		{Level(40.5), false},
		{Level(math.Inf(-1)), false},
		{Level(math.Inf(1)), false},
		{Level(math.NaN()), false},
	}
	for _, c := range cases {
		got := c.level.IsInteger()
		if got != c.expected {
			t.Errorf("Level(%f).IsInteger() = %t", c.level, got)
		}
	}
}

// TestSelf checks the method self() with predefined test cases.
func TestSelf(t *testing.T) {
	cases := []Level {
		Level(1.0),
		Level(1.5),
		Level(20.0),
		Level(20.5),
		Level(39.5),
		Level(40.0),
	}
	for _, c := range cases {
		got := c.self().id
		if got != c {
			t.Errorf("Level(%f).self().id = %f", c, got)
		}
	}
}
