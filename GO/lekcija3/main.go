package main

import "fmt"

// Ovako definiramo strukture(AKA model samo se tako ne ove)
type Person struct {
	firstName string
	lastName  string
	age       int
}

// Value reciver memtoda

func (p Person) greet() {
	fmt.Printf("hello my name is %s %s and I'm %d years old \n", p.firstName, p.lastName, p.age)
}

// Pointer reciver memtoda
func (p *Person) changeAge(age int) {
	p.age = age
}

type Books struct {
	name      string
	author    string
	year      int
	genre     string
	publisher Publisher
}
type Publisher struct {
	name    string
	address string
}

func main() {
	person := Person{
		firstName: "Blaž",
		lastName:  "Perč",
		age:       20,
	}
	fmt.Println(person)
	fmt.Println(person.firstName, person.lastName, person.age)
	person.greet()
	person.changeAge(25)
	person.greet()

	book := Books{
		name:   "Harry Potter",
		author: "J.K. Rowling",
		year:   1997,
		genre:  "Fantastika",
		publisher: Publisher{
			name:    "HarperCollins",
			address: "London"},
	}

	fmt.Println(book)
	fmt.Println("Ime knjige:", book.name)
	fmt.Println("Ime autora:", book.author)
	fmt.Println("Godina izdanja:", book.year)
	fmt.Println("Zanr:", book.genre)
	fmt.Println("Publisher name: ", book.publisher.name)
	fmt.Println("Publisher address: ", book.publisher.address)

}
