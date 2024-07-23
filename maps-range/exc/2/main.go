package main

func main() {
	mapinha := map[string]int{
		"manga":   10,
		"banana":  5,
		"uva":     15,
		"laranja": 20,
		"pera":    25,
	}

	valorTotal := getValorTotal(mapinha)
	println("Valor total:", valorTotal)
}

func getValorTotal(m map[string]int) int {
	t := 0
	for _, f := range m {
		t += f
	}
	return t
}
