package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(2)

	go print1()
	go print2()

	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}

func print1() {
	runtime.NumCPU()
	runtime.NumGoroutine()
	for i := 0; i <= 10; i++ {
		println(i)
		time.Sleep(50)
	}
	wg.Done()
}

func print2() {
	for i := 0; i <= 10; i++ {
		println(i)
		time.Sleep(50)
	}
	wg.Done()
}
