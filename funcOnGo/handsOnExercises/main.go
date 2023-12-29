package main

import (
	"fmt"
	"math"
)

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, ", and I am", p.age)
}

func main() {
	total := sum([]int{1, 2, 3, 4, 5, 6, 67, 7})
	total2 := prepareSum(1, 2, 3, 4, 5, 6, 67, 7)
	defer fmt.Println(total, total2)

	foo := foo()
	bar, barString := bar()
	defer fmt.Println(foo)
	fmt.Println(bar, barString)

	p := person{
		first: "Roger",
		age:   32,
	}

	defer p.speak()
	defer calcArea()
	mockDb()

	func() {
		fmt.Println("anonymous func")
	}()

	x := func() func() int {
		return func() int {
			return 1
		}
	}
	b := x()
	fmt.Println(b())

	c := doMath(1, 2, multNum)

	fmt.Println(c)

	d := Powinator(3)

	fmt.Println(d())
	fmt.Println(d())
	fmt.Println(d())
}

func sum(i []int) int {
	total := 0

	for _, v := range i {
		total += v
	}
	return total
}

func prepareSum(i ...int) int {
	return sum(i)
}

func foo() int {
	return 1
}

func multNum(a int, b int) int {
	return a * b
}

func bar() (int, string) {
	return 1, "True"
}

func doMath(a int, b int, op func(int, int) int) int {
	return op(a, b)
}

func Powinator(a float64) func() float64 {
	var c float64 = 0
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}
