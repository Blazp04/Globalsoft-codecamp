package model

import (
	"blazperic/lekcija8/zadatak1/errors"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) (err error) {

	if amount > b.Balance {
		err = &errors.WithdrawError{Owner: b.Owner, Amount: amount, Message: "Insufficient funds"}
		return
	}

	b.Balance -= amount
	return nil
}
