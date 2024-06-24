package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	Country string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func main() {
	fernando := Client{
		Name:   "Fernando",
		Age:    36,
		Active: true,
	}

	fernando.Age = 37

	fernando.Street = "Rua 1"
	fernando.Address.City = "Cordisburgo"

	fmt.Println(fernando.Name)

	fmt.Println(fernando.Address)
}
