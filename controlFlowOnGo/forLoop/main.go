package main

import "fmt"

var j = 0

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("counting to five %v\n", i)
	}

	for j < 10 {
		fmt.Printf("counting to ten %v\n", j)
		j++
	}

	for {
		v := 0

		if v > 20 {
			break
		}

		fmt.Printf("counting to 20 %v\n", v)
		v++
	}
}
