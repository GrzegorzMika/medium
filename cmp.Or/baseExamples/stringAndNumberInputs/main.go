package main

import "fmt"

func firstNonZeroString(words ...string) string {
	for _, word := range words {
		if word != "" {
			return word
		}
	}
	return ""
}

func firstNonZeroInt(ints ...int) int {
	for _, int := range ints {
		if int != 0 {
			return int
		}
	}
	return 0
}

func main() {
	strInputs := []string{"", "a", "b", "c", "d"}
	fmt.Println("Strings: ", firstNonZeroString(strInputs...))

	intInputs := []int{0, 1, 2, 3, 4}
	fmt.Println("Integers: ", firstNonZeroInt(intInputs...))
}
