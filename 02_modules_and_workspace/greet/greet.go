package greet

import "fmt"

// Hello returns a greeting for the given name.
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to Go modules.", name)
}
