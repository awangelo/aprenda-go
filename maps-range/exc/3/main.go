package main

import "fmt"

func main() {
	mapinha := map[string]int{"João": 85, "Maria": 52, "Pedro": 78, "Ana": 95, "Carlos": 68}

	media := getMedia(mapinha)
	fmt.Println("Média:", media)
}

func getMedia(m map[string]int) float64 {
	media := 0
	for _, n := range m {
		media += n
	}
	return float64(media / len(m))
}
