package main

import (
	"fmt"
	"sync"
)

func fanInOut() {
	eve := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	go sendFanInPut(eve, odd)

	receiveFanInPut(eve, odd, fanIn)

	for v := range fanIn {
		fmt.Printf("%v\t", v)
	}
}

func sendFanInPut(e, o chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func receiveFanInPut(e, o <-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			fanIn <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fanIn <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanIn)
}
