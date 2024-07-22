package main

import (
	"fmt"
	"time"
)

func main() {
	agora := time.Now()
	fmt.Printf("time.Now():\n%v\n\n", agora)

	// Formatado:
	fmt.Printf("time.Now().Format():\n%v\n\n", agora.Format("02/01/2006 15:04:05"))

	// Acessando os campos da data
	fmt.Println("Campos da data:")
	fmt.Println(agora.Year(), agora.Month(), agora.Day(), agora.Hour(), agora.Minute(), agora.Second())
	println()

	// Criando uma data
	minhaData := time.Date(1923, 07, 22, 15, 30, 0, 0, time.Local)
	fmt.Printf("time.Date():\n%v\n\n", minhaData)

}
