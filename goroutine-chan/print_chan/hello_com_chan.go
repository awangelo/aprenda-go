package main

import "fmt"

func main() {
	ch := make(chan string)
	go enviarOi(ch)

	fmt.Println(<-ch)
}

func enviarOi(c chan<- string) {
	minhaString := "Ola Alice!"
	c <- minhaString
}
