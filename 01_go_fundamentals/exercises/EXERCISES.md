# Module 01 — Exercises

Complete these after running `go run main.go` in the parent folder.

---

## Exercise 1: Temperature Converter

Create `exercises/ex1_temperature.go` with `package main` and write a function:

```go
func celsiusToFahrenheit(c float64) float64
```

Formula: `F = C * 9/5 + 32`

Test: `celsiusToFahrenheit(0)` should return `32`.

---

## Exercise 2: FizzBuzz

Write a function that prints numbers 1 to 20:
- Multiples of 3 → print "Fizz"
- Multiples of 5 → print "Buzz"
- Multiples of both → print "FizzBuzz"
- Otherwise print the number

---

## Exercise 3: Struct Practice

Create a `Book` struct with `Title`, `Author`, `Pages`.
Add a method `Summary() string` that returns `"Title by Author (Pages pages)"`.

---

## Exercise 4: Error Handling

Write `func sqrt(n float64) (float64, error)` that returns an error for negative numbers.

---

## How to Check Your Work

```bash
cd exercises
go run ex1_temperature.go   # after you create it
```

Or add tests in `exercises/exercises_test.go`.

---

## Solutions

Try exercises yourself first. Solutions are in `exercises/solutions/` (only peek after attempting).
