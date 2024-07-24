package main

import "fmt"

func main() {
	c := make(chan int)
	q := make(chan int)

	go infSelect(c, q)
	send(c, q)
}

func send(cs chan<- int, q chan int) {
	qlqrCoisa := 0
	for {
		select {
		case cs <- qlqrCoisa:
			qlqrCoisa++
		case <-q:
			fmt.Println("\nRecieved quit signal")
			return
		}
	}
}

func infSelect(cr <-chan int, q chan int) {
	for range 20 {
		fmt.Println("Recebido:", <-cr)
	}
	q <- -1
}
