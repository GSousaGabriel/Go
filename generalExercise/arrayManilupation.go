package main

import (
	"fmt"
)

func arrayM() {
	x := [5]int{}

	for i := 0; i < 5; i++ {
		x[i] = i
	}

	for i, v := range x {
		fmt.Printf("the value is %v and the index is %v\n", v, i)
	}
}
