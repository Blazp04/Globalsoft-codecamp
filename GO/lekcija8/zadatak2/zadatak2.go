package zadatak2

import (
	"fmt"

	"blazperic/lekcija8/zadatak2/model"
)

func Zadatak2() {
	book := model.Book{
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
