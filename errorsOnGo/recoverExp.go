package main

import "fmt"

func recoverExp() {
	test()
	fmt.Println("Normally printed")
}

func test() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in test", r)
		}
	}()
	test2(0)
	fmt.Println("Returned normally from test2")
}

func test2(i int) {
	if i > 3 {
		fmt.Println("Panicking!!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	test2(i + 1)
}
