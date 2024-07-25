package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("oi.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Fecha o arquivo apos a main terminar
	defer f.Close()

	// Retorna um slice de bytes (bs)
	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}
