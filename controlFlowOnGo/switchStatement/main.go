package main

import "fmt"

var x = 40

func main() {
	switch {
	case x < 45:
		fmt.Println("x is lower than needed")
	case x <= 40:
		fmt.Println("x is ideal")
	case x < 20:
		fmt.Println("x is too low")
	case x > 45:
		fmt.Println("x is higher than needed")
	default:
		fmt.Println("x is inadequate")
	}
	switch x {
	case 45:
		fmt.Println("x is 45")
	case 44:
		fmt.Println("x is 44")
	case 43:
		fmt.Println("x is 43")
	case 42:
		fmt.Println("x is 42")
	case 41:
		fmt.Println("x is 41")
	default:
		fmt.Println("x is default")
	}
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 44:
		fmt.Println("x is 44")
	case 43:
		fmt.Println("x is 43")
	case 42:
		fmt.Println("x is 42")
	case 41:
		fmt.Println("x is 41")
	default:
		fmt.Println("x is 40")
	}

}
