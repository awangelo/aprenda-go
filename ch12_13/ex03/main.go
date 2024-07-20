package main

import "fmt"

func main() {
	defer fmt.Println("Sou o primeiro print no codigo")
	defer fmt.Println("Sou o segundo print no codigo")
	fmt.Println("Sou o terceiro print no codigo")
	fmt.Println("Sou o quarto print no codigo")
}
