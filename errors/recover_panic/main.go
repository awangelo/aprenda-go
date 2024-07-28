package main

import (
	"fmt"
)

func main() {
	if err := safeFunc(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("okay ^^")
	}
}

func safeFunc() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()
	dangerousFunc()
	return nil
}

func dangerousFunc() {
	panic("deu ruim")
}
