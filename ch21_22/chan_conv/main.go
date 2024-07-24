package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

	go send(par, impar)
	go receive(par, impar, converge)
	// Converge tudo em um unico canal
	for v := range converge {
		fmt.Println("Recebido:", v)
	}
}

func send(p, i chan int) {
	for v := range 30 {
		if v%2 == 0 {
			p <- v
		} else {
			i <- v
		}
	}
	close(p)
	close(i)
}

func receive(p, i, c chan int) {
	var wg sync.WaitGroup
	defer close(c)
	defer wg.Wait()

	wg.Add(2)
	go func() {
		for v := range p {
			c <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range i {
			c <- v
		}
		wg.Done()
	}()
}
