package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(1, 3, 5, 7, 8, 12))
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
