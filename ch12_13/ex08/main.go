package main

import (
	"fmt"
)

func main() {
	a := getAscii('a')
	fmt.Print("Sua letra eh: ")
	a()
}

func getAscii(c rune) (f func()) {
	return func() {
		fmt.Printf("%U\n", c)
	}
}
