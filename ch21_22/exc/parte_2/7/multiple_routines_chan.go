package main

import "fmt"

const total int = 100

func main() {
	c := make(chan int)

	for range 10 {
		go send(c)
	}

	for range total {
		v := <-c
		fmt.Println("got:", v)
	}
}

func send(c chan<- int) {
	for v := range 10 {
		c <- v
	}
}
