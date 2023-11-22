package main

import (
	"fmt"
	"math/rand"
)

func loop() {
	x := 0

	// for i := 0; i <= 100; i++ {
	// 	fmt.Println(i)
	// }

	for x <= 42 {
		y := rand.Intn(5)

		switch y {
		case 1:
			fmt.Printf("Number %v for iteration %v\n", y, x)
		case 2:
			fmt.Printf("Number %v for iteration %v\n", y, x)
		case 3:
			fmt.Printf("Number %v for iteration %v\n", y, x)
		case 4:
			fmt.Printf("Number %v for iteration %v\n", y, x)
		default:
			fmt.Printf("Default %v for iteration %v\n", y, x)
		}
		x++
	}

	// for {
	// 	fmt.Println(x)
	// 	x++

	// 	if x == 100{
	// 		break
	// 	}
	// }
}
