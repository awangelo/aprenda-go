package main

import "fmt"

func main() {
	c := make(chan int)

	go MandarProCanal(c, 10)

	LerDoCanal(c)
}

func MandarProCanal(cs chan int, n int) {
	for i := 0; i < n; i++ {
		cs <- i
	}

	// Fechar o canal se nao ele vai ficar esperando por mais valores
	close(cs)
}

func LerDoCanal(cr <-chan int) {
	for i := range cr {
		fmt.Println("Recebido:", i+1)
	}
}
