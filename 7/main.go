package main

import "fmt"

func main() {
	salarios := map[string]int{"Fernando": 20000, "John": 19000, "Doe": 22000}

	fmt.Println(salarios["Fernando"])
	delete(salarios, "John")
	salarios["Fern"] = 23000
	fmt.Println(salarios)

	s := make(map[string]int)

	s["Fernando"] = 25000
	s["Doe"] = 30000

	fmt.Println(s)

	for name, salary := range s {
		fmt.Printf("Name: %s, Salary: %d\n", name, salary)
	}

	for _, salary := range s {
		fmt.Printf("Salary: %d\n", salary)
	}

}
