package main

import "fmt"

type MyNumbers interface {
	~int | float64
}

type MyAlias int

func main() {
	var n MyAlias = 42
	fmt.Println(addI(12, 2))
	fmt.Println(addF(1.41, 2.43))
	fmt.Println(addT(n, 2))
}

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

func addT[T MyNumbers](a, b T) T {
	return a + b
}
