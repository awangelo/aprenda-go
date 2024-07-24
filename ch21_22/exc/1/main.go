package main

import "fmt"

const enviar int = 20

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go send1(c1, enviar/2)
	go send2(c2, enviar/2)

	for range enviar {
		select {
		case v := <-c1:
			fmt.Println("Valor recebido no c1:", v)
		case v := <-c2:
			fmt.Println("Valor recebido no c2:", v)
		}
	}
}

func send1(s chan<- int, n int) {
	for i := range n {
		s <- i
	}
}

func send2(s chan<- int, n int) {
	for i := range n {
		s <- i
	}
}
