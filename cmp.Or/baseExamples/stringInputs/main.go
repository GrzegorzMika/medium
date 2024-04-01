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

func main() {
	inputs := []string{"", "a", "b", "c", "d"}
	fmt.Println("First example: ", firstNonZeroString(inputs...))

	inputs = []string{"", "", "b", "", "d"}
	fmt.Println("Second example: ", firstNonZeroString(inputs...))
}
