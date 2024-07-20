package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())

	var count int64 = 0
	totalGoRoutines := 10

	var wg sync.WaitGroup
	wg.Add(totalGoRoutines)

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			runtime.Gosched()
			fmt.Println("Count:", atomic.LoadInt64(&count))
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Count FINAL: ", count)
}
