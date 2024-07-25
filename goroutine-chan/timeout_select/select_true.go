package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	go func() {
		fmt.Println("sleeping...")
		time.Sleep(2 * time.Second) // Mudar este tempo
		ch <- true
	}()

	select {
	case b := <-ch:
		fmt.Println("Received!", b)
	case <-time.After(1 * time.Second): // Funciona como uma "flag" para executar o case apos o tempo
		fmt.Println("Timeout!")
	}
}
