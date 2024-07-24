package main

import "fmt"

type Animal interface {
	sound() string
	move() string
}

type Cachorro struct {
	Nome  string
	Idade int
}

func (c Cachorro) sound() string {
	return "rwuf"
}

func (c Cachorro) move() string {
	return "corre"
}

type Gato struct {
	Nome  string
	Idade int
}

func (g Gato) sound() string {
	return "meow"
}

func (g Gato) move() string {
	return "pula"
}

type Passaro struct {
	Nome  string
	Idade int
}

func (p Passaro) sound() string {
	return "piu"
}

func (p Passaro) move() string {
	return "voa"
}

func main() {
	a := []Animal{
		Cachorro{"Addie", 3},
		Gato{"Lily", 1},
		Passaro{"Pearl", 5},
	}

	for _, a := range a {
		fmt.Println(a.sound())
		fmt.Println(a.move())
		fmt.Println("")
	}
}
