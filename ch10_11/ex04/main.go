package main

import "fmt"

func main() {
	s := struct {
		cor      string
		ano      int
		ehZeroKm bool
	}{
		cor:      "preto",
		ano:      2021,
		ehZeroKm: true,
	}

	fmt.Println(s)
}
