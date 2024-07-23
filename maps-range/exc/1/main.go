package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "the cat jumped over the cat"
	freq := frequenciaPalavra(str)

	fmt.Println("Frequência de palavras")
	for k, v := range freq {
		fmt.Println("Palavra:", k, "\tFrequência:", v)
	}
}

func frequenciaPalavra(s string) map[string]int {
	freq := make(map[string]int)
	p := strings.Fields(s)

	for _, p := range p {
		freq[p]++
	}
	return freq
}
