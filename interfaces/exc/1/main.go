package main

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Circulo struct {
	Raio float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func imprimirArea(f Forma) {
	fmt.Printf("Area: %f\n", f.Area())
}

func main() {
	retangulo := Retangulo{Largura: 5, Altura: 3}
	circulo := Circulo{Raio: 2}

	imprimirArea(retangulo)
	imprimirArea(circulo)
}
