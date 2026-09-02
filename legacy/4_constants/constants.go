//  This program demonstrates the use of CONSTANTS in Go.
//  Constants are fixed values whose value cannot be changed
//  after they are defined.
//
//  Go constants are created using the `const` keyword.
//  They are evaluated at compile time and are useful for
//  values that should remain unchanged throughout the program
//  such as mathematical values, limits, configuration flags,
// and application-wide constants.

// package main

//  The fmt package is used to print output to the console.
// import "fmt"

// The const block allows defining multiple constants together.
//  This improves readability and keeps related values organized.
//
//  Go supports typed constants (int, float32, float64, bool, string)
//  and also supports untyped constants (not shown here).
//
//  Naming Rule (IMPORTANT):
//  - Constants starting with CAPITAL letters are exported
//    (accessible from other packages).
//  - Constants starting with lowercase letters are unexported
//    (accessible only within the same package).
// const (
// 	 PI is a floating-point constant representing the value of π.
// 	 It is exported because the name starts with a capital letter.
// 	PI float64 = 3.14159

// 	 name is a string constant that stores a person's name.
// 	 It is unexported because it starts with a lowercase letter.
// 	name string = "Suraj Rathod"

// 	IsEnabled is a boolean constant typically used as a feature flag
// 	 to enable or disable a specific functionality in the program.
// 	IsEnabled bool = true

// 	 ageLimit is an integer constant representing a minimum age requirement.
// 	 It is unexported and can only be accessed within this package.
// 	ageLimit int = 18

// 	 ageMin and ageMax define a valid age range.
// 	 These are useful for validation and boundary checking.
// 	ageMin int = 0
// 	ageMax int = 120

// 	 agelimit is a floating-point constant with a different data type (float32).
// 	 This shows that constants can be defined using different numeric types.
// 	agelimit float32 = 18.5

// 	 MaxAge is an exported floating-point constant.
// 	 It can be accessed from other packages if imported.
// 	MaxAge float32 = 100.0

// 	 Uncommenting the line below would add another constant,
// 	 but it is currently disabled to demonstrate comment usage.
// 	 ageMinf float32 = 0.0
// )

//  The main function is the entry point of the Go program.
//  Execution of the program starts from this function.
// func main() {

// 	 Printing constant values to the console using fmt.Println.
// 	 Constants behave like normal variables when printed,
// 	 but their values cannot be changed.
// 	fmt.Println("PI:", PI)
// 	fmt.Println("MaxAge:", MaxAge)
// 	fmt.Println("IsEnabled:", IsEnabled)
// 	fmt.Println("Name:", name)
// 	fmt.Println("Age Limit:", ageLimit)
// 	fmt.Println("Age Min:", ageMin)
// 	fmt.Println("Age Max:", ageMax)
// 	fmt.Println("Age Limit (float):", agelimit)

// 	 The commented line below will not execute.
// 	 It is shown here to demonstrate how comments can
// 	 disable code without deleting it.
// 	 fmt.Println("Age Min (float):", ageMinf)
// }

package main

import "fmt"

const (
	PI        float64 = 3.14159
	name      string  = "Suraj Rathod"
	IsEnabled bool    = true
	ageLimit  int     = 18   // unexported constant
	ageMin    int     = 0    // unexported constant
	ageMax    int     = 120  // unexported constant
	agelimit  float32 = 18.5 // unexported constant
	MaxAge    float32 = 100.0
	// ageMinf   float32 = 0.0 // unexported constant

)

func main() {
	fmt.Println("PI:", PI)
	fmt.Println("MaxAge:", MaxAge)
	fmt.Println("IsEnabled:", IsEnabled)
	fmt.Println("Name:", name)
	fmt.Println("Age Limit:", ageLimit)
	fmt.Println("Age Min:", ageMin)
	fmt.Println("Age Max:", ageMax)
	fmt.Println("Age Limit (float):", agelimit)
	// fmt.Println("Age Min (float):", ageMinf)
}
