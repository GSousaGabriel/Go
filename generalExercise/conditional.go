package main

import (
	"fmt"
	"math/rand"
)

func init() {
	x := rand.Intn(301)

	fmt.Printf("the value of x is %v\t\n", x)

	if x > 10 && x < 50 || x == 10 {
		fmt.Printf("the number %v is low\n", x)
	} else if x < 100 {
		fmt.Printf("the number %v is medium\n", x)
	} else if x <= 200 {
		fmt.Printf("the number %v is high\n", x)
	} else if x > 200 {
		fmt.Printf("the number %v is very high!\n", x)
	}

	switch {
	case x <= 100:
		fmt.Println("Too low")
	case x <= 250:
		fmt.Println("Almost there!")
	case x >= 250:
		fmt.Println("Allowed!")
	default:
		fmt.Println("Not Allowed")
	}
}
