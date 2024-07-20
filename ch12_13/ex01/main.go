package main

import "fmt"

func main() {
	fmt.Printf("O numero de `interio` eh: %v\n", inteiro())
	fmt.Printf(inteiroString())
}

func inteiro() int {
	return 1
}

func inteiroString() (int, string) {
	return 10, "O seu numero eh: "
}
