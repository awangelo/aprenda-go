package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	wg.Add(3)
	go enviarMsg(ch1, "Ola")
	go enviarMsg(ch2, "Alice")

	go func() {
		for {
			select {
			case resp, ok := <-ch1:
				if !ok {
					ch1 = nil
				} else {
					fmt.Println("Got a message from ch1:", resp)
				}
			case resp, ok := <-ch2:
				if !ok {
					ch2 = nil
				} else {
					fmt.Println("Got a message from ch2:", resp)
				}
			}
			if ch1 == nil && ch2 == nil {
				wg.Done()
				break
			}
		}
	}()
	wg.Wait()
	fmt.Println("end")
}

func enviarMsg(c chan<- string, m string) {
	defer wg.Done()
	c <- m
	close(c)
}
