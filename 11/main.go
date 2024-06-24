package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	fernando := Client{
		Name:   "Fernando",
		Age:    36,
		Active: true,
	}

	fernando.Age = 37

	fmt.Println(fernando.Name)
}
