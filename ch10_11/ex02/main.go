package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	saborFav  []string
}

func main() {
	mapinha := make(map[string]pessoa)

	mapinha["linddell"] = pessoa{"Alice", "Liddell", []string{"chocolate", "morango"}}
	mapinha["junior"] = pessoa{"Jhon", "Junior", []string{"baunilha", "creme"}}

	/*
		mapinha := map[string]pessoa{
		 	"linddell": {"Alice", "Liddell", []string{"chocolate", "morango"}},
		 	"junior":    {"Jhon", "Junior", []string{"baunilha", "creme"}},
		}
	*/

	for _, v := range mapinha {
		fmt.Println(v.nome)
		for _, v := range v.saborFav {
			fmt.Println("\t", v)
		}
	}

}
