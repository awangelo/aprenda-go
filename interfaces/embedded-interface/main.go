package main

import (
	"fmt"
)

type Animal interface {
	Pet
}

func printAnimal(a Animal) {
	fmt.Print("Nome do anima: ")
	println(a)
	fmt.Println("Som do animal:", a.Sound())
}

type printError struct {
	msg string
}

func (pr printError) Error() string {
	return pr.msg
}

func fazerBarulho(a Animal) (string, error) {
	if a.Sound() == "" {
		return "", printError{msg: "Som nao pode ser nulo"}
	}
	return a.Sound(), nil
}

type Pet interface {
	Sound() string
}

type Dog struct {
	Name string
}

func (d Dog) Sound() string {
	return "rwuf!" // Remover para testar o erro
}

type Cat struct {
	Name string
}

func (c Cat) Sound() string {
	return "meow"
}

func main() {
	dog := Dog{Name: "Rex"}
	cat := Cat{Name: "Felix"}

	fmt.Println("Usando interface")
	fmt.Println("Som do cachorro:", dog.Sound())
	fmt.Println("Som do gato:", cat.Sound())
	println()

	fmt.Println("Usando função com interface")
	printAnimal(dog)
	printAnimal(cat)
	println()

	fmt.Println("Usando função com interface e erro")
	som, err := fazerBarulho(dog)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Som do cachorro:", som)
	}
}
