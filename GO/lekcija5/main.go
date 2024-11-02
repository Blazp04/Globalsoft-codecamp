package main

import "fmt"

func main() {

	age := 12
	fmt.Println(checkAge(age))
	fmt.Println(switchChecking(age))

	fmt.Println(zadatak1(90))
	fmt.Println(zadatak1(80))
	fmt.Println(zadatak1(70))
	fmt.Println(zadatak1(60))
	fmt.Println(zadatak1(50))

}

func checkAge(age int) string {
	if age >= 60 {
		return "senior"
	}
	if age >= 18 {
		return "adult"
	}

	if age >= 13 {
		return "teenager"
	}
	return "child"
}

func switchChecking(age int) string {
	switch {
	case age >= 60:
		return "senior"
	case age >= 18:
		return "adult"
	case age >= 13:
		return "teenager"
	default:
		return "child"
	}
}

func zadatak1(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}
