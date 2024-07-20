package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracao4r bool
}
type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	carroDoTio := caminhonete{veiculo{2, "amarelo"}, true}
	carroDaVizinha := sedan{veiculo{4, "preto"}, true}

	fmt.Println(carroDoTio)
	fmt.Println(carroDaVizinha)
	println()

	fmt.Println("Do tio eh tracao nas 4 rodas?:", carroDoTio.tracao4r)
	fmt.Println("Da vizinha eh de luxo?:", carroDaVizinha.modeloLuxo)
}
