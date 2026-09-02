# 01. Go Fundamentals

> **Time:** 3–5 hours | **Prerequisites:** [Install Go](../docs/INSTALL.md)

Learn Go syntax from scratch — variables, control flow, structs, interfaces, and error handling.

---

## What You Will Learn

- [ ] Variables, zero values, and constants
- [ ] Control flow: `if`, `switch`, `for` (no `while` in Go!)
- [ ] Pointers — when and why
- [ ] Structs, methods, and interfaces (composition over inheritance)
- [ ] Error handling with `if err != nil`

---

## Step-by-Step Lessons

Follow `main.go` — each function is one lesson:

| Step | Function | Topic |
| :---: | :--- | :--- |
| 1 | `lesson01Variables()` | Variables, zero values, `const` |
| 2 | `lesson02ControlFlow()` | `if`, `switch`, `for` loops |
| 3 | `lesson03Pointers()` | `&` and `*` operators |
| 4 | `lesson04StructsAndInterfaces()` | Structs + implicit interfaces |
| 5 | `lesson05ErrorHandling()` | Functions returning `(value, error)` |

### Step 1 — Open and read the code

```bash
cd 01_go_fundamentals
# Open main.go in your editor — read lesson01Variables()
```

### Step 2 — Run the program

```bash
go run main.go
```

### Step 3 — Expected output

```
=== 01 Go Fundamentals ===

--- Lesson 1: Variables & Zero Values ---
Zero values → int: 0, string: "", bool: false
Hello, Go! | pi = 3.14

--- Lesson 2: Control Flow ---
Grade: B
Weekday: Monday
Sum 1..5 = 15

--- Lesson 3: Pointers ---
x=10, address=0x..., *ptr=10
After *ptr=20 → x=20

--- Lesson 4: Structs & Interfaces ---
Hi, I'm admin_user (ID: 1)

--- Lesson 5: Error Handling ---
10 / 2 = 5
Expected error: cannot divide 10 by zero
```

### Step 4 — Experiment

Change values in `main.go` and re-run. Try:
- Change `score` to 95 — what grade prints?
- Create a new `User` and pass it to `sayHello()`

---

## Exercises

Complete exercises in [exercises/EXERCISES.md](./exercises/EXERCISES.md):

1. Temperature converter
2. FizzBuzz
3. Book struct with method
4. `sqrt` with error handling

---

## Checkpoint

Before moving to Module 02, you should be able to:

- [ ] Explain zero values for `int`, `string`, `bool`
- [ ] Write a `for` loop and `if/else` without looking at docs
- [ ] Explain difference between `&` (address) and `*` (dereference)
- [ ] Create a struct and attach a method to it
- [ ] Handle errors without using `panic`

**Self-test:** Explain Go's approach to OOP in one sentence.
> *Answer hint: Composition via structs and implicit interfaces — no classes or inheritance.*

---

## Key Concepts

| Concept | Go Way |
| :--- | :--- |
| OOP | Structs + interfaces (implicit satisfaction) |
| Loops | Only `for` — no `while`, no `do-while` |
| Errors | Return `error` value — no try/catch |
| Unused imports | Compile error — forces clean code |

---

## Next Module

→ [02 Modules & Workspace](../02_modules_and_workspace/README.md)

Track progress: [docs/PROGRESS.md](../docs/PROGRESS.md)
