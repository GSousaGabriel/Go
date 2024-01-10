package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func atomicTest() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutine", runtime.NumGoroutine())

	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutine", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counter", counter)
}
