package main

import (
	"blazperic/lekcija8/zadatak1/model"
	"fmt"
)

func zadatak1() {
	anteB := model.BankAccount{
		Owner:   "Ante",
		Balance: 0,
	}
	anteB.Deposit(100)
	err := anteB.Withdraw(200)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Balance: %f\n", anteB.Balance)
}
