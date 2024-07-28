package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	if err := connectDB(); err != nil {
		fmt.Println(err)
	}
}

func connectDB() error {
	if err := connect(); err != nil {
		log.Println(err)
		return fmt.Errorf("Can't connect to the database: %w", err)
	}
	return nil
}

func connect() error {
	return errors.New("Database not found")
}
