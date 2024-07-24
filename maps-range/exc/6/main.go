package main

import "fmt"

func main() {
	msi := map[string]int{"banana": 3, "maca": 5, "pera": 2, "uva": 9, "cereja": 8}

	fmt.Println("Map antes:")
	fmt.Println("   key:value")
	fmt.Println(msi)
	println()

	mis := reverseMap(msi)
	fmt.Println("Map depois:")
	fmt.Println("   key:value")
	fmt.Println(mis)
}

func reverseMap(m map[string]int) map[int]string {
	new_map := make(map[int]string, len(m))

	for k, v := range m {
		new_map[v] = k
	}
	return new_map
}
