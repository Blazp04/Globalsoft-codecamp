package main

import "fmt"

func main() {
	zadatak4()
}

func zadatak4() {
	var total int = 150
	var discount float64 = 12.5
	finalPrice := float64(total) - (float64(total) * (discount / 100))

	fmt.Println(finalPrice)

}

func zadatak3() {
	var myIntegrer int
	var myFloat float64
	var myBool bool
	var myString string

	fmt.Println(myIntegrer, myFloat, myBool, myString)
}
func zadatak2() {
	const PI float64 = 3.14
	var radius float64 = 5.5

	opseg := 2 * PI * radius
	fmt.Println("Opseg kruga je:", opseg, "cm")

}
func zadatak1() {
	var firstName string = "Blaž Perić"
	var age int = 20
	var height float64 = 1.85
	var isStudent bool = true

	fmt.Println(firstName, age, height, isStudent)
}
