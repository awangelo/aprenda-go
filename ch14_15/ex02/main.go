package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func main() {
	a := pessoa{nome: "Alice", idade: 16}
	novoNome := "Angelo"
	novaIdade := 18

	fmt.Printf("Antes:  %v\n", a)
	mudarPessoa(&a, novoNome, novaIdade)
	fmt.Printf("Depois: %v\n", a)
}

func mudarPessoa(p *pessoa, s string, i int) {
	p.nome = s
	p.idade = i
}
