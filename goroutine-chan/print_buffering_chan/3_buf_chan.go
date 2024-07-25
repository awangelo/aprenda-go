package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	go enviarMsg(ch)

	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("main done ^^")
}

func enviarMsg(c chan<- string) {
	msg := []string{"Simple name.", "For a", "complicated being."}

	for _, m := range msg {
		c <- m
	}
	close(c)
}
