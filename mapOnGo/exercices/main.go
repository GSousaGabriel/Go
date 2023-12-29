package main

import "fmt"

func main() {
	mapFunc := map[string][]string{
		"bond_james":       {"shaken, not Stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
	}

	// mapFunc := make(map[string][]string)
	mapFunc["test"] = []string{"1", "2"}

	delete(mapFunc, "test")

	for k, v := range mapFunc {
		fmt.Println("key", k)
		fmt.Println("value", v)

		for key, val := range v {
			fmt.Println("key", key)
			fmt.Println("value", val)
		}
	}
}
