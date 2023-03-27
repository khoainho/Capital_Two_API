package main

import "math/rand"

type Account struct {
	ID int 
	FirstName string
	LastName string
	AcctNumber int64
	Balance int64
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