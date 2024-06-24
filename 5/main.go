package main

import "fmt"

type ID int

var (
	a         = "Hello, World!"
	b ID      = 2
	c float64 = 21.5845
)

func main() {
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	fmt.Println(myArray)
	fmt.Println(myArray[0])
	fmt.Println(len(myArray) - 1)

	for i, v := range myArray {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

}
