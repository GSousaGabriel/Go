package main

import "fmt"

func selectStat() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)

	receive(eve, odd, quit)
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Eve channel:", v)
		case v := <-o:
			fmt.Println("Odd channel:", v)
		case v := <-q:
			fmt.Println("Quit channel:", v)
			return
		}
	}
}
