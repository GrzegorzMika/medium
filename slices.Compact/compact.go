package main

import (
	"fmt"
	"slices"
)

func compactInternals() {
	a := []int{1, 1, 2, 3, 4}
	fmt.Printf("a = %d\n", a)
	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))
	// a = [1 1 2 3 4]
	// len(a) = 5
	// cap(a) = 5

	_ = slices.Compact(a)
	fmt.Printf("a = %d\n", a)
	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))
	// a = [1 2 3 4 0]
	// len(a) = 5
	// cap(a) = 5
}
