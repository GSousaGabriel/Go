package main

import "fmt"

func main() {
	mapFunc := map[string]int{
		"Tood":  42,
		"Roger": 56,
		"Henry": 31,
		"Mario": 2,
	}

	fmt.Printf("Age is: %v\n ", mapFunc["Henry"])

	an := make(map[string]int)

	an["Lucas"] = 28
	an["Maria"] = 346

	fmt.Printf("%v\n", an)
	delete(an, "Maria")

	fmt.Println("-----------------------------")

	for _, v := range an {
		fmt.Println(v)
	}

	if v, ok := an["mario"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("error")
	}

	fmt.Println("-----------------------------")

	a := 0
	fmt.Println(a)
	a++
	fmt.Println(a)

	m := make(map[string]int)

	fmt.Println(m["beautiful"])
}
