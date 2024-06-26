package main

import "fmt"

func main() {
	a := 10
	b := 2
	c := 3

	if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is not less than b")
	}

	if a > b || c > a {
		fmt.Println("a is less than b or c is greater than a")
	} else {
		fmt.Println("a is not less than b or c is not greater than a")
	}

	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	case 3:
		fmt.Println("a is 3")
	default:
		fmt.Println("a is not 1, 2 or 3")
	}
}
