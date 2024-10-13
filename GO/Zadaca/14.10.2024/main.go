package main

import "fmt"

func main() {
	zadatak1()
	zadatak2()
}

// Zadatak 1
// 1. Declare two integers, `firstNumber` and `secondNumber`, assign values 20 and 30 to them.
// Swap values of `firstNumber` and `secondNumber` without using third variable
// After all, print values of `firstNumber` and `secondNumber`.

func zadatak1() {
	var firstNumber int = 20
	var secondNumber int = 30

	firstNumber, secondNumber = secondNumber, firstNumber
	fmt.Println(firstNumber, secondNumber)
}

// Zadatak	2
// 2. Declare two variables, `firstName` and `lastName` assign them with wanted values.
// Declare constant named `fullname`
// Combine constant and both strings into a full name by concatenating strings with a space in between and print them out.
func zadatak2() {

	const firstName string = "Blaž"
	const lastName string = "Perićč"

	const fullname string = firstName + " " + lastName

	fmt.Println(fullname)
}
