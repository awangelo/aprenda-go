package main

import "fmt"

type pessoa struct {
	nome  string
	sobre string
	idade int
}

func main() {
	vizinha := pessoa{"Alice", "Liddell", 15}
	vizinha.printPessoa()
}

func (p pessoa) printPessoa() {
	fmt.Printf("Nome: %v\nSobrenome: %v\nIdade: %v\n", p.nome, p.sobre, p.idade)
}
