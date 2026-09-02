package main

import (
	"errors"
	"fmt"
)

type Calculator struct{}

func (c Calculator) Add(a, b int) int {
	return a + b
}

func (c Calculator) Multiply(a, b int) int {
	return a * b
}

func (c Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("=== 04 Testing & Debugging ===")
	fmt.Println("Run these commands:")
	fmt.Println("  go test -v          # run all tests")
	fmt.Println("  go test -bench=.    # run benchmarks")
	fmt.Println("  go test -fuzz=.     # run fuzz tests")
	fmt.Println("  go test -race .     # race detector")

	calc := Calculator{}
	fmt.Printf("Demo: 5 + 3 = %d\n", calc.Add(5, 3))
}
