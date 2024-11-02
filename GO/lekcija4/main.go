package main

import "fmt"

type Publisher struct {
	name           string
	location       string
	booksPublished int
}

type Book struct {
	title     string
	autor     string
	publisher Publisher
	price     float64
}

func (b Book) printDetails() {
	fmt.Printf("Title: %s\nAuthor: %s\nPublisher: %s\nPrice: %.2f\n ---------------------------\n", b.title, b.autor, b.publisher.name, b.price)
}

func (b *Book) applyDiscount(discountPercentage float64) {
	b.price = b.price * (1 - discountPercentage/100)
}

func (b *Book) updatePublisher(newPublisher string) {
	b.publisher.name = newPublisher
}

func (p *Publisher) printInfo() {
	fmt.Printf("Publisher name: %s\nPublisher location: %s\nNumber of books published: %d\n ---------------------------\n", p.name, p.location, p.booksPublished)
}
func (p *Publisher) publishBook() {
	p.booksPublished++
}

func (p *Publisher) changeLocation(newLocation string) {
	p.location = newLocation

}
func main() {
	publisher := Publisher{
		name:           "HarperCollins",
		location:       "London",
		booksPublished: 1,
	}

	book := Book{
		title:     "Harry Potter",
		autor:     "J.K. Rowling",
		publisher: publisher,
		price:     100,
	}

	book.printDetails()
	publisher.printInfo()

	book.applyDiscount(10)
	book.updatePublisher("Blaž Peirć")

	publisher.publishBook()
	publisher.changeLocation("Mostar")

	book.printDetails()
	publisher.printInfo()

}
