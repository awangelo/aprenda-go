package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	resp := make(chan int)

	go enviarNum(ch1, 12)
	go enviarNum(ch2, 13)
	go fanIn(resp, ch1, ch2)

	fmt.Println(<-resp)
	fmt.Println(<-resp)

	close(ch1)
	close(ch2)

	fmt.Println("done")
}

func enviarNum(c chan<- int, n int) {
	c <- n
}

func fanIn(resp chan int, c1, c2 <-chan int) {
	for {
		select {
		case n, ok := <-c1:
			if ok {
				resp <- n
			}
			c1 = nil
		case n, ok := <-c2:
			if ok {
				resp <- n
			}
			c1 = nil
		}
		if c1 == nil && c2 == nil {
			fmt.Println("closing fanIn")
			close(resp)
			break
		}
	}
}
