package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)

	go colocarNumero(ch, 3)
	for range 3 {
		wg.Add(1)
		go printNumero(ch)
	}
	wg.Wait()
	close(ch)

	fmt.Println("main done!")
}

func colocarNumero(c chan<- int, n int) {
	for v := range n {
		c <- v
	}
}

func printNumero(c <-chan int) {
	n := <-c
	fmt.Println("printNumero got:", n)
	wg.Done()
}
