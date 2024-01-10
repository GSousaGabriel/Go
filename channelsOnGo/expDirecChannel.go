package main

import "fmt"

func sendReceiveExp() {
	c := make(chan int)

	//send
	go foo(c)

	//receive
	for v := range c {
		fmt.Printf("%v\t", v)
	}

	//bar(c)
	fmt.Println("About to exit")
}

func foo(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
