package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func marshalData() {
	p1 := Person{
		Name: "Roger",
		Age:  14,
	}
	p2 := Person{
		Name: "Lucas",
		Age:  53,
	}

	people := []Person{p1, p2}
	fmt.Println(people)
	result, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(result))
}
