package main

import (
	"fmt"
	"runtime"
	"sync"
)

func mutex() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutine", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutine", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counter", counter)
}
