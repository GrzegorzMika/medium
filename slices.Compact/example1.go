package main

import (
	"fmt"
	"slices"
)

func example1() {
	a := []int{1, 1, 2, 3, 4}
	fmt.Printf("a = %d\n", a)
	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))
	// a = [1 1 2 3 4]
	// len(a) = 5
	// cap(a) = 5

	b := slices.Compact(a)
	fmt.Printf("b = %d\n", b)
	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("cap(b) = %d\n", cap(b))
	// b = [1 2 3 4]
	// len(b) = 4
	// cap(b) = 5
}
