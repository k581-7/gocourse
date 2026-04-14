package basics

import "fmt"

var middleName = "Garcia"

func main() {
	// var age int
	// var name string = "John"
	// var name1 = "Jane"

	// count := 10
	// lastName := "Smith"
	// middleName := "Garcia"

	// middleName := "Cane"
	fmt.Println(middleName)

	// DEFAULT VALUES
	// Numeric Types: 0
	// Boolean Types: false
	// String Type: ""
	// Pointers, slices, maps, functions, and structs: nil

	// ---- SCOPE

	// fmt.Println(firstName)

}

func printName() {
	firstName := "Karla"
	fmt.Println(firstName)
}
