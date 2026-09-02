# Module 05 — Exercises

## Exercise 1: Update User

Add `Update(user *User) error` to the repository.
Write SQL: `UPDATE users SET name=?, email=? WHERE id=?`

## Exercise 2: Delete User

Add `Delete(id int) error` and test deleting user ID 2.

## Exercise 3: Unique Constraint

Try inserting duplicate email — handle the SQLite unique constraint error gracefully.

## Exercise 4: File-Based SQLite

Change `:memory:` to `file:learning.db` — data persists between runs.

## Exercise 5: Repository Tests

Create `main_test.go` with table-driven tests for Create and GetByID using `:memory:` DB.
