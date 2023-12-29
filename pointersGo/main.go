package main

import "fmt"

type Dog struct {
	first string
}

func (d Dog) walk() {
	d.first = "Rover23"
	fmt.Println("Dog", d.first, "is walking")
}

func (d *Dog) run() {
	d.first = "Rover"
	fmt.Println("Dog", d.first, "is running")
}

type Youngin interface {
	walk()
	run()
}

func youngRun(y Youngin) {
	y.run()
}

func main() {
	x := 42

	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%v\t%T\n", &x, &x)

	y := &x
	fmt.Printf("%v\t%T\n", y, y)
	fmt.Println(*y)

	*y = 43
	fmt.Println(y)
	fmt.Println(x)

	a := 12
	fmt.Println(a)
	// defer fmt.Println(a) value not changed by intDelta
	intDelta(&a)
	fmt.Println(a)
	fmt.Println("--------------------------------------")
	b := []int{1, 2, 3, 4}
	fmt.Println(b)
	sliceDelta(b)
	fmt.Println(b)
	fmt.Println("--------------------------------------")

	d1 := Dog{"Joho"}
	d1.walk()
	fmt.Println(d1.first)
	d1.run()
	fmt.Println(d1.first)

	d2 := &Dog{"Roger"}
	fmt.Println(d2.first)
	youngRun(d2)
	fmt.Println(d2.first)
}

func intDelta(n *int) {
	fmt.Println("test")
	*n = 41
}

func sliceDelta(n []int) {
	fmt.Println("test2")
	n[0] = 32
}
