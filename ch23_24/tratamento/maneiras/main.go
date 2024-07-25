package main

import (
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("Aconteceu um erro:", err)
		//		log.Println("Aconteceu um erro:", err) // mostra a data e hora, pode enviar para um arquivo
		//		log.Fatalln(err) // `Println() followed by a call to os.Exit(1)`	encerra o programa e defer nao eh executado
		//		log.Panicln(err) // `Println() followed by a call to panic(err)`	encerra o programa e defer eh executado. 	PODE SER `RECOVER`
		//		panic(err) //
	}
}

// Descomentar as linhas para testar
// Recomendado usar os que tem `log.`
