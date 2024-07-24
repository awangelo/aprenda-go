package main

import "fmt"

type Forma interface {
	Volume() float64
	Superficie() float64
}

type Paralelepipedo struct {
	altura      float64
	largura     float64
	comprimento float64
}

func (c Paralelepipedo) Volume() float64 {
	return c.altura * c.largura * c.comprimento
}

func (c Paralelepipedo) Superficie() float64 {
	return 2 * (c.altura*c.largura + c.altura*c.comprimento + c.largura*c.comprimento)
}

type Esfera struct {
	raio float64
}

func (e Esfera) Volume() float64 {
	return 4 / 3 * 3.14 * e.raio * e.raio * e.raio
}

func (e Esfera) Superficie() float64 {
	return 4 * 3.14 * e.raio * e.raio
}

func main() {
	cubooo := Paralelepipedo{altura: 4, largura: 4, comprimento: 4}
	esferaa := Esfera{raio: 2}
	paralelele := Paralelepipedo{altura: 2, largura: 3, comprimento: 2}

	fmt.Println("Volume do cubo:", cubooo.Volume())
	fmt.Println("Superficie do cubo:", cubooo.Superficie())
	fmt.Println()

	fmt.Println("Volume da esfera:", esferaa.Volume())
	fmt.Println("Superficie da esfera:", esferaa.Superficie())
	fmt.Println()

	fmt.Println("Volume do paralelepipedo:", paralelele.Volume())
	fmt.Println("Superficie do paralelepipedo:", paralelele.Superficie())
}
