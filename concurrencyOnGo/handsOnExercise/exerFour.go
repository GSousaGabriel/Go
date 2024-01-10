package main

import (
	"fmt"
	"sync"
)

func rewriteCounterMutex() {
	var wg sync.WaitGroup

	counter := 0
	maxCount := 100
	wg.Add(maxCount)
	var mu sync.Mutex

	for i := 0; i < maxCount; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Print(counter, "\t")
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("\nCounter Mutex", counter)
}
