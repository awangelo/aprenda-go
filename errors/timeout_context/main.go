package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := performLongTask(ctx); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Long task done")
	}
}

func performLongTask(ctx context.Context) error {
	taskDuration := 1 * time.Second

	select {
	case <-ctx.Done():
		return fmt.Errorf("Ctx time out: %v", ctx.Err().Error())
	case <-time.After(taskDuration):
		fmt.Println("Task done!")
		return nil
	}
}
