package main

import (
	"fmt"
)

// Define a BankAccount struct with fields: `Owner` and `Balance`.
// Implement methods:
// - `Deposit`(amount): Adds money to the balance.
// - `Withdraw`(amount): Deducts money from the balance.
// If there are insufficient funds, return custom error with a message Insufficient funds.

// Write a main function to test depositing and withdrawing, checking for errors.

// -------------------------

// Define a Book struct with fields: `Title`, `Author`, and `Quantity`
// Create a function `AddBooks` that increases the Quantity of a Book by a given number.
// Write a function `RemoveBooks` that decreases the Quantity of a Book. If Quantity becomes negative, return custom error indicating Insufficient stocks.
// Implement a function `PrintBookInfo` that prints the details of the book, including available stock.

type BankAccount struct {
	owner   string
	balance float64
}

type WithdrawError struct {
	owner   string
	amaunt  float64
	message string
}

type Book struct {
	Title    string
	Author   string
	Quantity int
}

func (b *Book) PrintBookInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nQuantity: %d\n", b.Title, b.Author, b.Quantity)
}

func (b *Book) AddBooks(amount int) {
	b.Quantity += amount
}

func (b *Book) RemoveBooks(amount int) error {
	if b.Quantity-amount < 0 {
		return fmt.Errorf("insufficient books")
	}
	b.Quantity -= amount
	return nil
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

func main() {
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

	book := Book{
		Title:    "harry potter",
		Author:   "J.K. Rowling",
		Quantity: 10,
	}

	book.PrintBookInfo()
	book.AddBooks(5)
	book.PrintBookInfo()
	err1 := book.RemoveBooks(15)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	book.PrintBookInfo()

}
