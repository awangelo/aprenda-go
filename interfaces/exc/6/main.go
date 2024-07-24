package main

import "fmt"

type LeitorDeDados interface {
	Ler() (string, error)
}

func printLeitorDados(l LeitorDeDados) {
	c, err := l.Ler()
	if err != nil {
		fmt.Println("Erro ao ler dados:", err)
		return
	}
	fmt.Println(c)
}

type LeitorDeArquivo struct {
}

func (a LeitorDeArquivo) Ler() (string, error) {
	return "arquivo...", nil
}

type LeitorDeRede struct {
}

func (r LeitorDeRede) Ler() (string, error) {
	return "rede...", nil
}

func main() {
	printLeitorDados(LeitorDeArquivo{})
	printLeitorDados(LeitorDeRede{})
}
