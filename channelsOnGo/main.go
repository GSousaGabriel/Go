package main

import "fmt"

func main() {
	c := make(chan int)
	//insert into channel c <- n
	//c <- 42 error on runtime

	go func() {
		c <- 42
	}()
	//get from channel <- c
	fmt.Println(<-c)
	fmt.Println("----------------------------")

	d := make(chan int, 2)
	// d := make(<-chan int, 2) receive only
	// d := make(chan<- int, 2) send only
	d <- 53
	d <- 51
	fmt.Println(<-d)
	fmt.Println(<-d)
	fmt.Printf("%T\n", d)
	fmt.Println("----------------------------")
	sendReceiveExp()
	fmt.Println("----------------------------")
	selectStat()
	fmt.Println("----------------------------")
	test := make(chan int)
	go func() {
		test <- 1
		close(test)
	}()
	if v, ok := <-test; !ok {
		fmt.Println("Not allowed", ok)
	} else {
		fmt.Println("done", v)
	}
	fmt.Println("----------------------------")
	//fanInOut()
	fanOut()
}
