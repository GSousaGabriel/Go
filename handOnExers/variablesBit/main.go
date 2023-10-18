package main

import (
	"fmt"
	"math"
)

type Bytesize int

var y int

const (
	KB Bytesize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

const (
	_ = iota
	a
	b
	c
	d
	e
	f
)

func main() {
	exercise3()

	fmt.Println(math.Pi)
	bitwise()
	measure()
	exercise3()
	exercise4()
	exercise5()
}

func bitwise() {
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f)
}

func measure() {
	fmt.Printf("%b \t\t\t %b\n", KB, KB)
	fmt.Printf("%b \t\t\t %b\n", MB, MB)
	fmt.Printf("%b \t\t\t %b\n", GB, GB)
	fmt.Printf("%b \t\t\t %b\n", TB, TB)
	fmt.Printf("%b \t\t\t %b\n", TB, TB)
	fmt.Printf("%b \t\t\t %b\n", EB, EB)
}

func exercise3() {
	fmt.Println(y)
	z := 42
	fmt.Println(z)

	a, b := 43, "hapiness"
	fmt.Println(a, b)

	var c float32 = 42.42
	fmt.Printf("%v is of this type %T\n", c, c)

	e, f, _ := 45, 46, 47

	fmt.Println(e, f)
}

func exercise4() {
	x, y, z := 747, 911, 90210

	fmt.Printf("%d \t\t %b \t\t %#x\n", x, x, x)
	fmt.Printf("%d \t\t %b \t\t %#x\n", y, y, y)
	fmt.Printf("%d \t\t %b \t\t %#x\n", z, z, z)
}

func exercise5() {
	var i int8 = 127
	var j uint8 = 255

	fmt.Printf("%v is of type %T\n", i, i)
	fmt.Printf("%v is of type %T\n", j, j)
}
