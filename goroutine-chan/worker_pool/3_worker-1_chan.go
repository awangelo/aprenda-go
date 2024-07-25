package main

import (
	"fmt"
	"sync"
)

const nWorkers int = 3

var wg sync.WaitGroup

func main() {
	// Nao precisa de buffer!
	// Mas se for enviar pela main, precisa de buffer (3)
	tarefas := make(chan int)

	go func() {
		defer close(tarefas)
		tarefas <- 3
		tarefas <- 12
		tarefas <- 1213998899
	}()

	for range nWorkers {
		wg.Add(1)
		go worker(tarefas)
	}
	wg.Wait()

	fmt.Println("main done ^w^")
}

func worker(c <-chan int) {
	defer wg.Done()

	numero := <-c
	fmt.Printf("Quadrado de %v eh: %v\n", numero, numero*numero)
}
