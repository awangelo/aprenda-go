package main

import "fmt"

func main() {
	m := map[string]int{
		"salgado": 7,
		"bolo":    30,
		"doce":    3,
	}

	for key, value := range m {
		fmt.Println("Key:", key, "\tValue:", value)
	}
	println()

	fmt.Println("Ignorando a chave")
	for key := range m {
		fmt.Println("Chave:", key)
	}
	println()

	fmt.Println("Ignorando o valor")
	for _, value := range m {
		fmt.Println("Valor:", value)
	}
	println()

	fmt.Println("Verificando se a chave existe")
	v, ok := m["bolo"]
	fmt.Println("Existe?", ok)
	fmt.Println("Valor:", v)
	v1, ok := m["chocolate"]
	fmt.Println("Existe?", ok)
	fmt.Println("Valor:", v1)
}
