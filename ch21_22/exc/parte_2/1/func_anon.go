package main

import (
	"fmt"
)

func main() {
	// com buffer:	c := make(chan int, 1)

	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
