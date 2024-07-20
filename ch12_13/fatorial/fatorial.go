package main

import "fmt"

func main() {
	var meuNumero uint64 = 50
	fmt.Printf("O fatorial de %v eh: %v\n", meuNumero, fact(meuNumero))
}

func fact(x uint64) uint64 {
	if x <= 1 {
		return x
	}
	return x * fact(x-1)
}
