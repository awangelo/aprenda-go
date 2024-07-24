package main

import "fmt"

func main() {
	c := make(chan int)

	go recieveQuit(c)
	sendQuit(c)
}

func recieveQuit(c chan int) {
	for {
		select {
		case <-c:
			fmt.Println("Recebendo 02")
		}
	}
}

func sendQuit(c chan int) {
	fmt.Println("Enviando 02")
	c <- 0
}
