package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var x, y = 40, 2

	if x < 50 {
		fmt.Println("x is less than the value asked")
	}
	if x < 50 && y > 3 {
		fmt.Println("y is a better choice than x")
	} else {
		fmt.Printf("both values are too low, since x is %v and y is %v\n", x, y)
	}

	if z := 3 * rand.Intn(x); z >= x {
		fmt.Printf("value x is tested as %v of type %T for a z equal to %v\n", x, x, z)
	} else {
		fmt.Printf("missed the shot x was %v and z was %v\n", x, z)
	}
}
