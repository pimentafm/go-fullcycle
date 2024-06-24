package main

import "fmt"

type ID int

var (
	a         = "Hello, World!"
	b ID      = 2
	c float64 = 21.5845
)

func main() {
	fmt.Printf("O tipo de a é %T\n", a)
	fmt.Printf("O tipo de b é %T\n", b)
	fmt.Printf("O tipo de c é %T\n", c)
}
