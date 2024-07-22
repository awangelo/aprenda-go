package main

import (
	"fmt"
	"time"
)

func main() {
	dataRecebida := "22/09/2024 15:30:00 -03"

	dataParsed, err := time.Parse("02/01/2006 15:04:05 -07", dataRecebida)
	if err != nil {
		panic(err)
	}

	agora := time.Now()
	fmt.Println("Agora:", agora)
	fmt.Println("Data formatada:", dataParsed)
	println()

	// Comparando as datas
	fmt.Println("Agora antes da formatada?", agora.Before(dataParsed))
	fmt.Println("Agora depois da formatada?", agora.After(dataParsed))
	fmt.Println("Agora igual a formatada?", agora.Equal(dataParsed))

}
