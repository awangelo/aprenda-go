package main

import "fmt"

func main() {
	var resposta1 string

	// CTRL+D = panic: EOF
	n, err := fmt.Scan(&resposta1)
	if err != nil {
		panic(err)
	}

	fmt.Println(resposta1, "\trecebido: ", n, "items")
}
