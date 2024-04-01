package main

import (
	"cmp"
	"fmt"
)

func main() {
	strInputs := []string{"", "a", "b", "c", "d"}
	fmt.Println("Strings: ", cmp.Or(strInputs...))

	intInputs := []int{0, 1, 2, 3, 4}
	fmt.Println("Integers: ", cmp.Or(intInputs...))

	floatInputs := []float64{0, 1.2, 0.171, 243, -1.23}
	fmt.Println("Floats: ", cmp.Or(floatInputs...))
}
