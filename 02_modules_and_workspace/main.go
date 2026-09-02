package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/suraj-iot-engineer/golang-learning/02_modules_and_workspace/greet"
)

func main() {
	color.Cyan("=== 02 Modules & Workspace ===")

	// Lesson 1: External dependency (from go.mod)
	c := color.New(color.FgHiGreen, color.Bold)
	c.Println("✅ External package: github.com/fatih/color")

	// Lesson 2: Local module (greet package in same workspace)
	fmt.Println(greet.Hello("Learner"))

	// Lesson 3: Standard library
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))

	lesson04Commands()
}

func lesson04Commands() {
	fmt.Println("\n--- Lesson 4: Essential Commands ---")
	fmt.Println("go mod init <module>  → create new module")
	fmt.Println("go get <package>      → add/update dependency")
	fmt.Println("go mod tidy           → remove unused deps")
	fmt.Println("go work init          → create workspace (multi-module)")
	fmt.Println("go work use ./greet   → add module to workspace")
}
