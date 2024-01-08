package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func waitGroupPractice() {
	wg.Add(2)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("foo:", i)
		}
		wg.Done()
	}()
	go bar()
	wg.Wait()
}

func bar() {
	for i := 0; i < 4; i++ {
		fmt.Println("bar:", i)
	}
	wg.Done()
}
