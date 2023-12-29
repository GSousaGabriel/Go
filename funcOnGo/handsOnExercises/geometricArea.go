package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

// calculate and returns the area of a square using math package
func (s square) area() float64 {
	area := s.width * s.length
	return area
}

// calculate and returns the area of a circle using math package
func (c circle) area() float64 {
	area := math.Pi * math.Pow(c.radius, 2)
	return area
}

func calcArea() {
	c := circle{
		radius: 423.3,
	}
	defer fmt.Println(c.area())

	s := square{
		length: 32.32,
		width:  43.1,
	}
	fmt.Println(s.area())
}
