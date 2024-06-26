package main

import "fmt"

type Wallet struct {
	balance float64
}

func (c *Wallet) simulate(value float64) float64 {
	c.balance += value

	return c.balance
}

func main() {
	wallet := Wallet{balance: 100.0}

	wallet.simulate(200)

	fmt.Printf("Balance simulation %.2f\n", wallet.balance)

}
