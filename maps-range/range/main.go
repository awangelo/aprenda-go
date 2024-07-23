package main

import "fmt"

func main() {
	si := []int{107, 238, 531, 235, 853}
	msi := map[string]int{"banana": 3, "maca": 5, "pera": 2, "uva": 9, "cereja": 8}

	fmt.Println("Range com um valor em um slice:")
	for i := range si {
		fmt.Println(i)
	}
	println()

	fmt.Println("Range com um valor em um map:")
	for v := range msi {
		fmt.Println(v)
	}
}
