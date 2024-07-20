package main

import (
	"fmt"
	"sort"
)

func main() {
	si := []int{42, 17, 89, 63, 5, 76, 31, 54, 10, 92, 27, 71, 49, 14, 98}

	sort.Ints(si)

	fmt.Println("Sorted:", si)
}
