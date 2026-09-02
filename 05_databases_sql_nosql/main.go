package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type User struct {
	ID    int
	Name  string
	Email string
}

type UserRepository interface {
	GetByID(id int) (*User, error)
	Create(user *User) error
	List() ([]User, error)
}

// SQLiteRepository uses real database/sql with SQLite (in-memory).
type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) (*SQLiteRepository, error) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL
		)
	`)
	if err != nil {
		return nil, err
	}
	return &SQLiteRepository{db: db}, nil
}

func (r *SQLiteRepository) Create(user *User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (id, name, email) VALUES (?, ?, ?)",
		user.ID, user.Name, user.Email,
	)
	return err
}

func (r *SQLiteRepository) GetByID(id int) (*User, error) {
	row := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
	u := &User{}
	if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *SQLiteRepository) List() ([]User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func main() {
	fmt.Println("=== 05 Databases (SQL) ===")

	// Lesson 1: Open database connection
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Lesson 2: Connection pool settings (important in production)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	repo, err := NewSQLiteRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	// Lesson 3: CRUD with parameterized queries (prevents SQL injection)
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
	}
	for _, u := range users {
		if err := repo.Create(&u); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("✅ Created: %s\n", u.Name)
	}

	// Lesson 4: Read single row
	found, err := repo.GetByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("🔍 Found: %+v\n", found)

	// Lesson 5: List all rows
	all, err := repo.List()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("📋 Total users: %d\n", len(all))

	// Lesson 6: Handle not found
	_, err = repo.GetByID(99)
	if err == sql.ErrNoRows {
		fmt.Println("✅ Correctly handled: user 99 not found")
	}
}
