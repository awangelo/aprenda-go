package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	if n, err := divide(2, 2); err != nil {
		fmt.Println("Error dividing:", err)
	} else {
		fmt.Println("Result:", n)
	}

	result, err := divide(4, 0)
	if err != nil {
		log.Fatalln("Error on division:", err)
	}
	fmt.Println("Result", result)
}

func divide(n, d int) (int, error) {
	if d == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return n / d, nil
}
