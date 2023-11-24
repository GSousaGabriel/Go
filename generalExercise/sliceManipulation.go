package main

import (
	"fmt"
)

func sliceM() {
	x := []int{}
	x = append(x, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51)

	for i, v := range x {
		fmt.Printf("the value is %v and the index is %v\n", v, i)
	}

	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

	y := []int{12, 1, 4, 14}
	x = append(x, y...)

	fmt.Println(x)

	z := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	z = append(z[:3], z[6:]...)
	fmt.Println(z)

	h := []string{"James", "BOND", "Not tired"}
	j := []string{"Miss", "Money", "I'm 008"}
	hj := [][]string{h, j}
	fmt.Println(hj)
}
