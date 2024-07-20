package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	saborFav  []string
}

func main() {
	alice := pessoa{"Alice", "Liddell", []string{"chocolate", "morango"}}
	jhon := pessoa{"Jhon", "Mama", []string{"baunilha", "creme"}}

	fmt.Println(alice)
	fmt.Println(jhon)
	println()

	fmt.Println(alice.nome, alice.sobrenome)
	for _, s := range alice.saborFav {
		fmt.Println(s)
	}
	println()
	fmt.Println(jhon.nome, jhon.sobrenome)
	for _, s := range jhon.saborFav {
		fmt.Println(s)
	}

}
