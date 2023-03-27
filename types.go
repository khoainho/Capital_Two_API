package main

import "math/rand"

type Account struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	AcctNumber int64 `json:"acct_number"`
	Balance int64 `json:"current_Balance"`
}

// Make temp account
func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID: rand.Intn(10000),
		FirstName: firstName,
		LastName: lastName,
		AcctNumber: int64(rand.Intn(1000000)),
	}
}