package main

import (
	"fmt"
	"time"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go send(par, impar, quit)
	receive(par, impar, quit)
}

func send(p, i chan<- int, q chan<- bool) {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range numeros {
		time.Sleep(200 * time.Millisecond)
		if v%2 == 0 {
			p <- v
		} else {
			i <- v
		}
	}
	close(p)
	close(i)
	q <- true
}

func receive(p, i <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-p:
			fmt.Println("Recebi um par:", v)
		case v := <-i:
			fmt.Println("Recebi um impar:", v)
		case <-q:
			fmt.Println("Recebi um quit sig.")
			return
		}
	}
}
