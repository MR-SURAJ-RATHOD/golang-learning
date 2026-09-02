// only constructs used to demonstrate 'for' loops in Go
// This program demonstrates all possible forms of `for` loops in Go.
// Go has ONLY ONE looping construct: `for`.
// There is no `while` or `do-while` loop in Go.
// Different loop behaviors are achieved by changing the syntax of `for`.

// package main

// The fmt package is used for formatted output to the console.
// It provides functions like Println and Printf for displaying text.
// import "fmt"

// The main function is the entry point of the Go program.
// Execution of the program always starts from main().
// func main() {

// ----------------------------------------------------
// SIMPLE FOR LOOP (CLASSIC COUNTER LOOP)
// ----------------------------------------------------
// This loop initializes a variable `i` to 0,
// runs as long as `i` is less than 10,
// and increments `i` by 1 after each iteration.
// This type of loop is used when the number of iterations is known.
// fmt.Println("Simple for loop:")
// for i := 0; i < 10; i++ {
// 	fmt.Println("Iteration:", i)
// }

// ----------------------------------------------------
// FOR LOOP WITH MULTIPLE VARIABLES
// ----------------------------------------------------
// This loop demonstrates how multiple variables can be
// initialized and updated inside a single for loop.
// Here:
// - `i` starts from 0 and increases by 1
// - `j` starts from 10 and decreases by 1
// The loop continues until `i` is less than `j`.
// This pattern is commonly used in algorithms such as
// array reversal and two-pointer techniques.
// fmt.Println("\nFor loop with multiple variables:")
// for i, j := 0, 10; i < j; i, j = i+1, j-1 {
// 	fmt.Printf("i: %d, j: %d\n", i, j)
// }

// ----------------------------------------------------
// FOR LOOP USED AS A WHILE LOOP
// ----------------------------------------------------
// Go does not have a `while` loop.
// A `for` loop with only a condition behaves like `while`.
// The loop continues until the condition becomes false.
// This is useful for condition-based iteration.
// fmt.Println("\nFor loop as a while loop:")
// count := 0
// for count < 6 {
// 	fmt.Println("Count:", count)
// 	count++
// }

// ----------------------------------------------------
// INFINITE FOR LOOP WITH BREAK
// ----------------------------------------------------
// A `for {}` loop creates an infinite loop.
// It runs forever unless explicitly stopped using `break` or `return`.
// Infinite loops are very common in:
// - Servers
// - IoT applications
// - Background workers
// Here, the loop exits when `n` becomes 3.
// 	fmt.Println("\nInfinite for loop with break:")
// 	n := 0
// 	for {
// 		if n >= 3 {
// 			break // Terminates the loop immediately
// 		}
// 		fmt.Println("n:", n)
// 		n++
// 	}
// }

package main

import "fmt"

func main() {

	// Simple for loop
	fmt.Println("Simple for loop:")
	for i := 0; i < 10; i++ {
		fmt.Println("Iteration:", i)
	}

	// For loop with multiple variables
	fmt.Println("\nFor loop with multiple variables:")
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Printf("i: %d, j: %d\n", i, j)
	}

	// For loop as a while loop
	fmt.Println("\nFor loop as a while loop:")
	count := 0
	for count < 6 {
		fmt.Println("Count:", count)
		count++
	}

	// Infinite for loop with break
	fmt.Println("\nInfinite for loop with break:")
	n := 0
	for {
		if n >= 3 {
			break
		}
		fmt.Println("n:", n)
		n++
	}
}
