package main

import (
	"fmt"
)

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

func zadatak2() {
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
