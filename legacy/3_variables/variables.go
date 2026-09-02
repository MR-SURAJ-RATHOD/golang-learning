// 3_variables/variables.go declares multiple variables of different types and prints their values.

// This program demonstrates how to declare and use VARIABLES in Go.
// Variables are used to store data that can change during
// program execution.
//
// Go is a statically typed language, which means every variable
// has a specific data type (int, float64, string, bool, etc.).
// The type of a variable determines what kind of data it can hold
// and what operations can be performed on it.

// package main

// The fmt package provides functions for formatted output,
// such as Println, which is used to display values on the console.
// import "fmt"

// The main function is the entry point of the program.
// All executable code must be written inside a function.
// func main() {

	// ----------------------------------------------------
	// VARIABLE DECLARATION USING A VAR BLOCK
	// ----------------------------------------------------
	// The var block allows declaring multiple variables together.
	// Each variable has:
	// - a name
	// - a data type
	// - an initial value
	//
	// This approach improves readability and keeps related
	// variables organized in one place.
	// var (
		// age stores an integer value representing a person's age.
		// age int = 25

		// height stores a floating-point value.
		// float64 is commonly used for decimal values
		// because it provides higher precision.
		// height float64 = 5.9

		// name stores text data.
		// The string type is used for words and sentences.
		// name string = "Suraj Rathod"

		// isStudent stores a boolean value.
		// Boolean variables can have only two values:
		// true or false.
		// isStudent bool = false

		// nickName stores an additional string value.
		// Variables can be used to store multiple related
		// pieces of information about an entity.
		// nickName string = "handsome"
	)

	// ----------------------------------------------------
	// PRINTING VARIABLE VALUES
	// ----------------------------------------------------
	// fmt.Println prints text and variable values to the console.
	// Each call prints the output on a new line.
// 	fmt.Println("Name:", name)
// 	fmt.Println("Age:", age)
// 	fmt.Println("Height:", height)
// 	fmt.Println("Is Student:", isStudent)
// 	fmt.Println("NickName:", nickName)
// }



package main

import "fmt"

func main() {
	// Declaring multiple variables of different types
	var (
		age       int     = 25
		height    float64 = 5.9
		name      string  = "Suraj Rathod"
		isStudent bool    = false
		nickName  string  = "handsome"
	)
	// Printing the variable values
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("NickName:", nickName)
}
