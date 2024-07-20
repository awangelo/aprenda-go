package main

import (
	"fmt"
)

func main() {
	c := make(chan int) // Create an unbuffered channel of type int

	go func() {
		c <- 42 // Send the value 42 to the channel c
	}()

	fmt.Println(<-c) // Receive the value from the channel c and print it
}
