package main

import "fmt"

func main() {
	c, err := fmt.Println("Ola Alice")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)

	c, err = fmt.Println("Ola Alice")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)

	// ???
}
