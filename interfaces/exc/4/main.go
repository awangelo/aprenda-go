package main

import "fmt"

type OperacaoMatematica interface {
	Calcular(a, b int) int
}

func fazerConta(o OperacaoMatematica, a, b int) int {
	return o.Calcular(a, b)
}

type Soma struct {
}

func (s Soma) Calcular(a, b int) int {
	return a + b
}

type Multiplicacao struct {
}

func (m Multiplicacao) Calcular(a, b int) int {
	return a * b
}

func main() {
	a, b := 10, 3
	soma := Soma{}
	multiplicacao := Multiplicacao{}

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println()

	fmt.Println("Calculando:")
	fmt.Println("Soma:", fazerConta(soma, a, b))
	fmt.Println("Multip.:", fazerConta(multiplicacao, a, b))
}
