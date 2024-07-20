package main

import (
	"fmt"
)

func main() {
	meusNumeros := []int{12, 34, 23, 7, 41, 18, 29, 5, 16, 38}
	res := impares(soma, meusNumeros...)
	fmt.Println(meusNumeros)
	fmt.Printf("Resultado da soma dos impares: %v\n", res)
}

func impares(f func(x ...int) int, y ...int) int {
	var impares []int
	for _, v := range y {
		if v%2 != 0 {
			impares = append(impares, v)
		}
	}
	return f(impares...)
}

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}
