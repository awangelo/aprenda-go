package main

import "fmt"

func main() {
	meusNumerooos := []int{5, 2, 7, 11, 5}
	fmt.Printf("Soma variatica: %v\n", somarVariatico(meusNumerooos...))
	fmt.Printf("Soma slices variatica: %v\n", somarSlicesVariatico(meusNumerooos))
}

func somarVariatico(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}

func somarSlicesVariatico(si []int) int {
	total := 0
	for _, v := range si {
		total += v
	}
	return total
}
