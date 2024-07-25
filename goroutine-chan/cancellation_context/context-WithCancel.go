package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go printSleep(ctx)

	time.Sleep(3 * time.Second)
}

func printSleep(ctx context.Context) {
	for i := range 10 {
		select {
		case <-ctx.Done():
			fmt.Println("printSleep done ^w^")
			return
		default:
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}
}
