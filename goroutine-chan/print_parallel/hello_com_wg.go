package main

import (
	"fmt"
	"sync"
)

const numeroDeLoops int = 5

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go hello(&wg, numeroDeLoops)
	wg.Wait()

	fmt.Println("main done ^^")
}

func hello(wg *sync.WaitGroup, n int) {
	defer wg.Done()
	for range n {
		fmt.Println("Hello from goroutine!")
	}
}
