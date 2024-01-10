package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := Person{
		First: "Miss",
		Last:  "Money",
		Age:   22,
	}
	people := []Person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}
	jsonString := string(bs)

	fmt.Println(jsonString)

	err = json.Unmarshal(bs, &jsonString)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All data", people)

	for i, v := range people {
		fmt.Println("Person number", i)
		fmt.Println(v.First)
		fmt.Println(v)
	}
}
