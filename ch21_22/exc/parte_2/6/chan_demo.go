package main

import "fmt"

func main() {
	c := make(chan int)

	go send(c)

	for v := range c {
		fmt.Println("recebi:", v)
	}
	fmt.Println("fim")
}

func send(c chan<- int) {
	for v := range 100 {
		c <- v
	}
	close(c)
}
