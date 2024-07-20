package main

import "fmt"

type humano interface {
	falar()
}

type pessoa struct {
	nome   string
	idade  int
	altura float64
	vivo   bool
}

func main() {
	a := pessoa{"Alice", 17, 1.70, true}
	j := pessoa{"Joe", 25, 1.80, true}

	dizerAlgo(&a)
	dizerAlgo(&j)
}

func (p *pessoa) falar() {
	fmt.Printf("Ola, sou uma pessoa (%v)\n", p.nome)
	fmt.Println("Tambem sou um metodo implementado por um tipo `Pessoa`")
}

func dizerAlgo(h humano) {
	h.falar()
}
