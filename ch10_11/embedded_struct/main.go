package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

type estudante struct {
	pessoa
	curso string
}

func main() {
	p1 := pessoa{"Alice", 20}
	fmt.Println(p1)
	p2 := pessoa{"Bob", 30}
	fmt.Println(p2)
	println()

	e1 := estudante{p1, "Letras"}
	fmt.Println(e1)

	println()
	fmt.Println(e1.nome)
	fmt.Println(e1.idade)
	fmt.Println(e1.curso)
}
