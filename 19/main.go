package main

import "fmt"

type MyNumber int

type Number interface {
	~int | float64
}

func Sum[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}

	return sum
}

func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Fernando": 1000, "Pimenta": 500}
	m2 := map[string]MyNumber{"Fernando": 1000, "Pimenta": 500}
	fmt.Println(Sum(m))
	fmt.Println(Sum(m2))

	fmt.Println(Compare(10, 10.0))
}
