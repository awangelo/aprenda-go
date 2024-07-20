package main

import (
	"fmt"
	"math"
)

type figura interface {
	getArea() float64
}

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func main() {
	meuQuadrado := quadrado{lado: 2}
	meuCirculo := circulo{raio: 3}

	fmt.Println("Área do quadrado:", info(meuQuadrado))
	fmt.Println("Área do círculo:", info(meuCirculo))
}

func (q quadrado) getArea() float64 {
	return q.lado * q.lado
}

func (c circulo) getArea() float64 {
	return math.Pi * c.raio * c.raio
}

func info(f figura) float64 {
	return f.getArea()
}
