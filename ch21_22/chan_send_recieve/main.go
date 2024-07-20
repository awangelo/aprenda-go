package main

import "fmt"

func main() {
	// apenasRecebe := make(chan<- int) // Create a receive-only channel of type int
	// apenasEnvia := make(<-chan int) // Create a send-only channel of type int

	c := make(chan int)

	go send(c, 1231)

	// Se for uma goroutine nao funciona, pois a main termina antes
	recieve(c)
}

func send(s chan<- int, v int) {
	s <- v
}

func recieve(s <-chan int) {
	fmt.Println("O valor recebido foi:", <-s)
}
