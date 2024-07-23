package main

import "fmt"

func main() {
	var meu_map map[string]int

	mapa1 := map[string]int{
		"salgado": 7,
		"bolo":    30,
		"doce":    3,
	}

	mapa2 := make(map[string]int, 10)

	fmt.Println("Meu Mapa:", meu_map)
	fmt.Println("Mapa 1:", mapa1)
	fmt.Println("Mapa 2:", mapa2)
	fmt.Println()

	fmt.Printf("Acessando o valor na chave '%s': %d\n", "chave1", mapa1["chave1"])
	println()

	fmt.Println("Colocando um novo par: 'mapa1[almoco] = 1213'")
	fmt.Printf("Antes: %v\n", mapa1)
	mapa1["almoco"] = 12
	fmt.Printf("Depois: %v\n\n", mapa1)

	fmt.Println("Deletando um par: 'delete(mapa1, almoco)'")
	fmt.Printf("Antes: %v\n", mapa1)
	delete(mapa1, "almoco")
	fmt.Printf("Depois: %v\n\n", mapa1)

	fmt.Println("Verificando se a chave 'almoco' existe no mapa1")
	valor, ok := mapa1["almoco"]
	fmt.Printf("Valor: %d, Existe: %v\n\n", valor, ok)

	fmt.Println("Verificando se a chave 'doce' existe no mapa1")
	valor, ok = mapa1["doce"]
	fmt.Printf("Valor: %d, Existe: %v\n\n", valor, ok)

	fmt.Println("Tamanho do mapa1:", len(mapa1))

}
