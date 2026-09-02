package main

import (
	"errors"
	"fmt"
	"time"
)

// === CORE / DOMAIN LAYER ===

type Order struct {
	ID        string
	Amount    float64
	CreatedAt time.Time
}

// OrderRepository is a PORT (Output Port)
type OrderRepository interface {
	Save(order Order) error
}

// OrderService is the Application Logic
type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(id string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	order := Order{
		ID:        id,
		Amount:    amount,
		CreatedAt: time.Now(),
	}

	return s.repo.Save(order)
}

// === ADAPTERS LAYER ===

// InMemoryOrderRepo is an ADAPTER (Secondary Adapter)
type InMemoryOrderRepo struct {
	store map[string]Order
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{store: make(map[string]Order)}
}

func (r *InMemoryOrderRepo) Save(order Order) error {
	r.store[order.ID] = order
	fmt.Printf("💾 Database: Saved order %s with amount %.2f\n", order.ID, order.Amount)
	return nil
}

// CLIHandler is an ADAPTER (Primary Adapter)
type CLIHandler struct {
	service *OrderService
}

func (h *CLIHandler) HandleCreateOrder(id string, amount float64) {
	err := h.service.CreateOrder(id, amount)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ Order created successfully")
	}
}

func main() {
	fmt.Println("=== 09 Backend Architecture (Hexagonal) ===")

	// 1. Initialize Adapters (Infrastructure)
	repo := NewInMemoryOrderRepo()

	// 2. Initialize Core (Domain) - Dependency Injection
	service := NewOrderService(repo)

	// 3. Initialize Primary Adapter (EntryPoint)
	handler := &CLIHandler{service: service}

	// 4. Execute
	handler.HandleCreateOrder("ORD-101", 99.99)
	handler.HandleCreateOrder("ORD-102", -10.00) // Should fail
}
