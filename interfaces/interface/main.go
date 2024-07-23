package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func getArea(s Shape) float64 {
	return s.Area()
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	circulo := Circle{radius: 2.5}
	retangulo := Rectangle{width: 4, height: 3}

	fmt.Println("Usando interface")
	fmt.Println("Área do círculo:", circulo.Area())
	fmt.Println("Área do retângulo:", retangulo.Area())
	println()

	fmt.Println("Usando função com interface")
	fmt.Println("Área do círculo:", getArea(circulo))
	fmt.Println("Área do retângulo:", getArea(retangulo))
}
