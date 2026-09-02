package main

import (
	"fmt"
	"log"
)

// User demonstrates structs and methods.
type User struct {
	ID       int
	Username string
	Active   bool
}

// Greeter is implemented by User (implicit interface).
type Greeter interface {
	Greet() string
}

func (u User) Greet() string {
	return fmt.Sprintf("Hi, I'm %s (ID: %d)", u.Username, u.ID)
}

func main() {
	fmt.Println("=== 01 Go Fundamentals ===")

	lesson01Variables()
	lesson02ControlFlow()
	lesson03Pointers()
	lesson04StructsAndInterfaces()
	lesson05ErrorHandling()
}

// --- Lesson 1: Variables & Zero Values ---
func lesson01Variables() {
	fmt.Println("\n--- Lesson 1: Variables & Zero Values ---")

	var count int
	var name string
	var active bool
	fmt.Printf("Zero values → int: %d, string: %q, bool: %v\n", count, name, active)

	message := "Hello, Go!" // short declaration
	const pi = 3.14
	fmt.Println(message, "| pi =", pi)
}

// --- Lesson 2: Control Flow ---
func lesson02ControlFlow() {
	fmt.Println("\n--- Lesson 2: Control Flow ---")

	// if (no parentheses needed)
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	// switch
	day := "Monday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Weekday:", day)
	}

	// for loop (Go has no while — use for)
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Printf("Sum 1..5 = %d\n", sum)
}

// --- Lesson 3: Pointers ---
func lesson03Pointers() {
	fmt.Println("\n--- Lesson 3: Pointers ---")

	x := 10
	ptr := &x
	fmt.Printf("x=%d, address=%p, *ptr=%d\n", x, ptr, *ptr)

	*ptr = 20 // change value through pointer
	fmt.Printf("After *ptr=20 → x=%d\n", x)
}

// --- Lesson 4: Structs & Interfaces ---
func lesson04StructsAndInterfaces() {
	fmt.Println("\n--- Lesson 4: Structs & Interfaces ---")

	admin := User{ID: 1, Username: "admin_user", Active: true}
	sayHello(admin) // User satisfies Greeter interface
}

func sayHello(g Greeter) {
	fmt.Println(g.Greet())
}

// --- Lesson 5: Error Handling ---
func lesson05ErrorHandling() {
	fmt.Println("\n--- Lesson 5: Error Handling ---")

	if res, err := divide(10, 2); err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %d\n", res)
	}

	if _, err := divide(10, 0); err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", a)
	}
	return a / b, nil
}
