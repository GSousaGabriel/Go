package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()

	fmt.Println(s1, "test", s2)
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
}
