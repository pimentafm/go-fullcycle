package main

import "fmt"

func sum(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	a := 10
	fmt.Println(a, &a)

	var pointer *int = &a

	fmt.Println(pointer, *pointer)

	*pointer = 20

	fmt.Println(a, &a)

	b := &a
	*b = 30

	fmt.Println(b, *b, a)

	var1 := 10
	var2 := 20

	sum(&var1, &var2)

	fmt.Println(var1, var2)
}
