package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)

	// Numero de goroutines (workers), cada uma retornara um valor
	funções := 5

	// 100 eh o numero de valores que serao enviados para o canal
	// ou seja, 100 valores serao enviados para o canal1 de 5 em 5 (5 goroutines)
	go manda(100, canal1)

	go outra(funções, canal1, canal2)
	for v := range canal2 {
		fmt.Println(v)
	}
}

func manda(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func outra(funções int, canal1, canal2 chan int) {
	var wg sync.WaitGroup
	for i := 0; i < funções; i++ {
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalho(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(time.Second)
	return n
}
