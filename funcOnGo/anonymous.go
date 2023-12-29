package main

import "fmt"

func anMain() {
	foo()

	func() {
		fmt.Println("anonyous ran")
	}()

	x := func(s string) {
		fmt.Println("Name is", s)
	}

	x("Terry")

	h := doMath(51, 42, add)
	fmt.Println(h)

	i := doMath(13, 312, substract)
	fmt.Println(i)

	j := testClosure()
	fmt.Println(j())
	fmt.Println(j())
	fmt.Println(j())
	fmt.Println(j())
}

func foo() {
	fmt.Println("Foo ran")
}

func doMath(a int, b int, cb func(int, int) int) int {
	return cb(a, b)
}

func add(a int, b int) int {
	return a + b
}

func substract(a int, b int) int {
	return a - b
}

func testClosure() func() int {
	x := 0
	fmt.Println("ran")
	return func() int {
		x++
		return x
	}
}
