package main

import "fmt"

func main() {
	c := make(chan int)

	go recieveQuit(c)
	go sendQuit(c)
}

func recieveQuit(c chan int) {
	for i := 0; i < count; i++ {

	}
}

func sendQuit(c chan int) {
	fmt.Println("Enviando 02")
}
