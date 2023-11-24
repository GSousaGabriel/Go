package main

import (
	"fmt"
	"math/rand"
)

func main2() {
	// x := []int{42, 43, 44, 45, 46, 47}
	mapVar := map[string]int{
		"james": 10,
		"test":  12,
	}

	for i, v := range mapVar {
		fmt.Printf("index %v, value %v\n", i, v)
	}

	age := mapVar["james"]
	fmt.Printf("James age is %v\n", age)

	age = mapVar["Q"]

	if v, ok := mapVar["Q"]; !ok {
		fmt.Printf("There is no Q %v\n", v)
	}

	fmt.Println("------------------")

	for i := 0; i < 10; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Println("x is", x)
		}
	}
}
