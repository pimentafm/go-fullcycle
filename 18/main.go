package main

import "fmt"

func main() {
	var myVar interface{} = "Fernando Pimenta"
	println(myVar.(string))
	res, ok := myVar.(int)

	fmt.Printf("res: %v, ok: %v\n", res, ok)

	res2 := myVar.(int)

	fmt.Printf("res2: %v\n", res2)
}
