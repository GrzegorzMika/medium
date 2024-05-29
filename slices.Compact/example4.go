package main

import (
	"fmt"
	"slices"
)

func sum(number ...int) int {
	sum := 0
	for _, n := range number {
		sum += n
	}
	return sum
}

func example4() {
	type Container struct {
		entries []int
	}

	items := []Container{
		{entries: []int{1, 5, 9}},
		{entries: []int{3, 4, 8}},
		{entries: []int{7, 8, 3}},
	}

	fmt.Printf("items: %v\n", items)
	// items: [{[1 5 9]} {[3 4 8]} {[7 8 3]}]

	newItems := slices.CompactFunc(items,
		func(item1, item2 Container) bool { return sum(item1.entries...) == sum(item2.entries...) })
	fmt.Printf("newItems: %v\n", newItems)
	// newItems: [{[1 5 9]} {[7 8 3]}]
}
