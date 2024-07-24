package main

import "fmt"

func main() {
	m := map[string]string{"rosa": "vermelho", "girassol": "amarelo", "violeta": "roxo", "margarida": "branco", "tulipa": "vermelho"}

	mapCat(m)
}

func mapCat(m map[string]string) {
	ks := ""
	vs := ""
	for k, v := range m {
		ks += k + " "
		vs += v + " "
	}
	fmt.Println(ks)
	fmt.Println(vs)
}
