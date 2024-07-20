package main

import "fmt"

func main() {
	x := retornaFunc(10)
	y := x(20)
	fmt.Println(y)
}

func retornaFunc(x int) func(x int) int {
	fmt.Println("Olha essa funcao:")
	return func(y int) int {
		return x + y
	}
}
