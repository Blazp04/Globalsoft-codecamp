package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet())
	names := []string{"Blaž", "Ante"}
	poydrav(names...)
	person := Person{
		firstName: "Blaž",
		lastName:  "Perć",
		age:       20,
	}
	name, lastName := pozzStukturi(person)

	fmt.Println(name, lastName)

	r1, e1 := zadatak1(2, 5, "+")
	fmt.Println(r1, e1)

	r2, e2 := zadatak1(2, 5, "-")
	fmt.Println(r2, e2)

	r3, e3 := zadatak1(2, 5, "*")
	fmt.Println(r3, e3)

	r4, e4 := zadatak1(2, 5, "/")
	fmt.Println(r4, e4)

	_, e5 := zadatak1(2, 5, "%")
	if e5 != "" {
		fmt.Println(e5)
	}
}

func greet() string {
	return "hello"
}

func poydrav(names ...string) {
	for _, name := range names {
		fmt.Printf("hello my name is %s \n", name)
	}

}

type Person struct {
	firstName string
	lastName  string
	age       int
}

func pozzStukturi(person Person) (name string, lastName string) {

	name = person.firstName
	lastName = person.lastName

	return
}

func zadatak1(n1 int, n2 int, operator string) (int, string) {

	if operator == "+" {
		return n1 + n2, ""
	}
	if operator == "-" {
		return n1 - n2, ""
	}
	if operator == "*" {
		return n1 * n2, ""
	}
	if operator == "/" {
		return n1 / n2, ""
	}
	return 0, "Niste unijeli ispravan operator"
}
