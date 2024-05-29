package main

import (
	"fmt"
	"slices"
)

func example2() {
	a := []int{1, 2, 1, 2, 1, 2}
	fmt.Printf("a = %d\n", a)
	// a = [1 2 1 2 1 2]

	b := slices.Compact(a)
	fmt.Printf("b = %d\n", b)
	// b = [1 2 1 2 1 2]
}
