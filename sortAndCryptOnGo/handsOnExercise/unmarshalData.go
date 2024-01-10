package main

import (
	"encoding/json"
	"fmt"
)

func unmarshalData() {
	stringJson := `[{"Name": "Gord", "Age": 23},{"Name": "Roger", "Age": 32}]`

	var people []Person

	err := json.Unmarshal([]byte(stringJson), &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for i, person := range people {
		fmt.Println("Person on position", i)
		fmt.Println(person.Name, person.Age)
		fmt.Println(person)
	}
}
