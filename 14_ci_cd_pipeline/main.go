package main

import "fmt"

func main() {
	fmt.Println("=== 14 CI/CD Pipeline (Local) ===")
	fmt.Println("Run .\\build_and_test.ps1 from repo root to test all modules!")
}

func Add(a, b int) int {
	return a + b
}
