package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"n"`
	Balance int `json:"b" validate:"gte=0"`
}

func main() {
	account := Account{Number: 1, Balance: 100}
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}

	rawJson := []byte(`{"n":2,"b":200}`)
	var account2 Account

	err = json.Unmarshal(rawJson, &account2)
	if err != nil {
		panic(err)
	}

	println(account2.Balance)
}
