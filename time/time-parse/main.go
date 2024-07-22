package main

import (
	"fmt"
	"time"
)

func main() {
	dataRecebida := "22/07/2024 15:30:00 -03"

	dataParsed, err := time.Parse("02/01/2006 15:04:05 -07", dataRecebida)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data recebida:", dataRecebida)
	fmt.Println("Data formatada:", dataParsed)

}
