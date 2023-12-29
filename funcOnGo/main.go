package main

import (
	"fmt"
	"log"
	"strconv"
)

type person struct {
	first string
}
type book struct {
	title string
}

type count int

func main() {
	defer fmt.Println("run last --------------------------------")
	s1, n := yourDogYears("Gabriel", 12)

	fmt.Println(s1, n)
	fmt.Println("variatic --------------------------------")

	total := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(total)

	numbers := []int{1, 2, 3, 4, 5, 7}
	total2 := sum(numbers...)
	fmt.Println(total2)

	p1 := person{
		first: "Roger",
	}

	p2 := person{
		first: "Rogeria",
	}

	p1.speak()
	p2.speak()

	fmt.Println("wrapper --------------------------------")

	b := book{
		title: "Tester",
	}

	var c count = 52

	logInfo(b)
	logInfo(c)

	writerTest()
	anMain()
}

func yourDogYears(name string, age int) (string, int) {
	age *= 7

	return fmt.Sprint(name, ", in dog years, you are"), age
}

func sum(i ...int) int {
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	n := 0

	for _, number := range i {
		n += number
	}

	return n
}

func (p person) speak() {
	fmt.Println("I am ", p.first)
}

func (b book) String() string {
	return fmt.Sprint("The title is ", b.title)
}

func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM 133 ", s.String())
}
