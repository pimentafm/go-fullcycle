package main

import (
	"errors"
	"fmt"
)

func main() {
	value, error := sum(5, 3)

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(value)
}

func sum(a int, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior ou igual a 50")
	}

	return a + b, nil
}
