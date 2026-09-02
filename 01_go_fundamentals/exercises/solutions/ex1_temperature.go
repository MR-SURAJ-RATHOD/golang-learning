package main

import "fmt"

func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func main() {
	fmt.Printf("0°C = %.0f°F\n", celsiusToFahrenheit(0))
	fmt.Printf("100°C = %.0f°F\n", celsiusToFahrenheit(100))
}
