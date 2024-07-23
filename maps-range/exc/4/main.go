package main

import "fmt"

func main() {
	m := map[string]int{
		"um":     1,
		"dois":   2,
		"tres":   3,
		"quatro": 4,
		"cinco":  5,
		"seis":   6,
		"sete":   7,
		"oito":   8,
		"nove":   9,
		"dez":    10,
	}

	removerMenorQue(m, 5)

	for k, v := range m {
		fmt.Println("Chave:", k, "\tValor:", v)
	}
}

func removerMenorQue(m map[string]int, n int) {
	for k, v := range m {
		if v < n {
			delete(m, k)
		}
	}
}
