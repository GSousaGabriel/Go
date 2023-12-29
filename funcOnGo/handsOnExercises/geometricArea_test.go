package main

import (
	"math"
	"testing"
)

func TestCalcArea(t *testing.T) {
	c := circle{
		radius: 1,
	}

	s := square{
		length: 2,
		width:  2,
	}

	areaS := s.area()
	areaC := c.area()

	if areaC != math.Pi {
		t.Errorf("Circle area is wrong, got: %f, want %f.", areaC, math.Pi)
	}

	if areaS != float64(4) {
		t.Errorf("Circle area is wrong, got: %f, want %f.", areaS, float64(4))
	}
}
