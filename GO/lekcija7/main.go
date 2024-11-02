package main

import (
	"fmt"
)

// error handling

type divisionError struct {
	a       int
	b       int
	message string
}

func (d *divisionError) Error() string {
	return fmt.Sprintf("Cannot devide %d by %d. %s", d.a, d.b, d.message)
}
func main() {

	a := 10
	b := 2

	result, err := devide(a, b)

	if err != nil {
		// %v je univerzalni olaceholder
		fmt.Printf("Error while deviding 2 numbers %v", err.Error())
		return
	}
	fmt.Println(result)
}

func devide(a int, b int) (int, error) {
	// Uvijek je dobro da dodamo pokazivač u reciver da možemo imati null vrijednost jer  je to jedini način da
	// vratimo null
	// Zato u gore reciver funckiji treba je definirati da passamo pokazivač(pointer) a ne vrijednost
	if b == 0 {
		return 0, &divisionError{
			a:       a,
			b:       b,
			message: "cannot devide by zero"}
	}
	return a / b, nil
}
