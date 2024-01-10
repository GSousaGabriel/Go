package main

import (
	"fmt"
	"runtime"
	"sync"
)

func rewriteCounter() {
	var wg sync.WaitGroup

	counter := 0
	maxCount := 100
	wg.Add(maxCount)

	for i := 0; i < maxCount; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Print(counter, "\t")
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("\nCounter", counter)
}
