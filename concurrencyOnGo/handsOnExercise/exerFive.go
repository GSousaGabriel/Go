package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func rewriteCounterAtomic() {
	var wg sync.WaitGroup

	var counter int64
	maxCount := 100
	wg.Add(maxCount)

	for i := 0; i < maxCount; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Print(atomic.LoadInt64(&counter), "\t")
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("\nCounter Atomic", counter)
}
