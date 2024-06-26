package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	numbers := []string{"one", "two", "three", "four", "five"}

	for k, v := range numbers {
		fmt.Println(k, v)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println("Infinite loop")
	}
}
