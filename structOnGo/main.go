package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	last string
	ltk  bool
}

func main() {
	//anon struct
	ps := struct {
		test bool
		work string
	}{
		test: true,
		work: "don't know",
	}

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   54,
	}

	sa1 := secretAgent{
		person: p1,
		last:   "yes",
		ltk:    true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Gabriel",
			last:  "Bond",
			age:   32,
		},
		last: "roger",
		ltk:  true,
	}

	fmt.Println(ps)

	fmt.Println(sa1.last, sa1.person.last)

	fmt.Println(sa2.last, sa2.ltk)
}
