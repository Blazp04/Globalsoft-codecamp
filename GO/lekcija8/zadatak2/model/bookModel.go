package model

import "fmt"

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
