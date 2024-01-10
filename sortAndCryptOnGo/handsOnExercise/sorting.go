package main

import (
	"fmt"
	"sort"
)

func sortingData() {
	intData := []int{1, 54, 2, 4, 867, 12, 64, 67235, 123, 1}
	stringData := []string{"Roger", "James", "Amanda", "Silva", "Carlos", "Mathias"}

	sort.Ints(intData)
	sort.Strings(stringData)
	fmt.Println(intData)
	fmt.Println(stringData)
}
