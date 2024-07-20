package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines ANTES:", runtime.NumGoroutine())

	count := 0
	totalGoRoutines := 10

	var wg sync.WaitGroup
	wg.Add(totalGoRoutines)
	var mu sync.Mutex

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			v := count
			runtime.Gosched()
			v++
			count = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines LOOP:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("GoRoutines FINAL:", runtime.NumGoroutine())
	fmt.Println("Count:", count)
}
