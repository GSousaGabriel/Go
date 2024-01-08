package main

import "fmt"

type Person struct {
	name string
}

type Human interface {
	speak()
}

func (p *Person) speak() {
	fmt.Println(p.name, "is talking")
}

func methodSets() {
	p := Person{
		name: "Roger",
	}

	saySomething(&p)
}

func saySomething(h Human) {
	h.speak()
}
