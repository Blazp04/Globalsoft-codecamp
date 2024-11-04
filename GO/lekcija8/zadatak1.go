package main

import (
	"fmt"
)

type BankAccount struct {
	owner   string
	balance float64
}

type WithdrawError struct {
	owner   string
	amaunt  float64
	message string
}

func (w *WithdrawError) Error() string {
	return fmt.Sprintf("Cannot withdraw %f from %s. %s", w.amaunt, w.owner, w.message)
}

func (b *BankAccount) deposit(amount float64) {
	b.balance += amount
}

func (b *BankAccount) withdraw(amount float64) (err error) {

	if amount > b.balance {
		err = &WithdrawError{owner: b.owner, amaunt: amount, message: "Insufficient funds"}
		return
	}

	b.balance -= amount
	return nil
}

func zadatak1() {
	anteB := BankAccount{
		owner:   "Ante",
		balance: 0,
	}
	anteB.deposit(100)
	err := anteB.withdraw(200)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Balance: %f\n", anteB.balance)
}
