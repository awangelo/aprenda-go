package main

import "fmt"

func main() {
	mp := map[string]float64{
		"arroz":    10.0,
		"feijao":   5.0,
		"macarrao": 3.0,
		"carne":    20.0,
		"frango":   15.0,
		"batata":   2.0,
	}

	alterarValorPorc(mp, 10.0)
	fmt.Println("Valores alterados:")
	for k, v := range mp {
		fmt.Println("Chave:", k, "\tValor:", v)
	}
}

func alterarValorPorc(m map[string]float64, x float64) {
	for k, v := range m {
		m[k] = v * (1 + x/100)
	}
}
