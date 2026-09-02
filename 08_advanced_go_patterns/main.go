package main

import (
	"fmt"
	"reflect"
	"time"
)

// --- Lesson 1: Generics ---

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(val T) { s.elements = append(s.elements, val) }

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	i := len(s.elements) - 1
	val := s.elements[i]
	s.elements = s.elements[:i]
	return val, true
}

// --- Lesson 2: Functional Options ---

type Server struct {
	Host    string
	Port    int
	Timeout time.Duration
}

type ServerOption func(*Server)

func NewServer(opts ...ServerOption) *Server {
	s := &Server{Host: "localhost", Port: 8080, Timeout: 30 * time.Second}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithPort(port int) ServerOption {
	return func(s *Server) { s.Port = port }
}

func WithTimeout(d time.Duration) ServerOption {
	return func(s *Server) { s.Timeout = d }
}

// --- Lesson 3: Middleware Chain ---

type Handler func(string) string

func chain(h Handler, middlewares ...func(Handler) Handler) Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func loggingMiddleware(next Handler) Handler {
	return func(s string) string {
		fmt.Println("  [log] processing request")
		result := next(s)
		fmt.Println("  [log] done")
		return result
	}
}

func uppercaseMiddleware(next Handler) Handler {
	return func(s string) string {
		return next(s + " (processed)")
	}
}

// --- Lesson 4: Reflection (use sparingly) ---

func inspectStruct(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Printf("  Type: %s, Kind: %s, Fields: %d\n", t.Name(), t.Kind(), t.NumField())
}

func main() {
	fmt.Println("=== 08 Advanced Patterns ===")

	fmt.Println("\n--- Lesson 1: Generics ---")
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	val, _ := intStack.Pop()
	fmt.Printf("  Stack pop: %d\n", val)

	fmt.Println("\n--- Lesson 2: Functional Options ---")
	srv := NewServer(WithPort(9000), WithTimeout(time.Minute))
	fmt.Printf("  Server: %+v\n", srv)

	fmt.Println("\n--- Lesson 3: Middleware Chain ---")
	handler := chain(
		func(s string) string { return s },
		loggingMiddleware,
		uppercaseMiddleware,
	)
	fmt.Printf("  Result: %s\n", handler("hello"))

	fmt.Println("\n--- Lesson 4: Reflection ---")
	inspectStruct(Server{Host: "0.0.0.0", Port: 8080})
	fmt.Println("  ⚠️ Prefer generics/interfaces over reflection in production")
}
