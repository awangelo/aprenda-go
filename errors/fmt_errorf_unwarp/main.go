package main

import (
	"errors"
	"fmt"
)

func main() {
	if err := getData(); err != nil {
		fmt.Println("Erro:", err)
		fmt.Println("Erro original:", errors.Unwrap(err))
	}
}

func fetchSomeData() error {
	originalError := errors.New("failed to connect to server")
	return fmt.Errorf("Cannot fetchData: %w", originalError)
}

func getData() error {
	return fetchSomeData()
}
