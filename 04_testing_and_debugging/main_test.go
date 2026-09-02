package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	calc := Calculator{}
	if got := calc.Add(2, 3); got != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", got)
	}
}

func TestMultiply(t *testing.T) {
	calc := Calculator{}
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive", 3, 4, 12},
		{"zero", 5, 0, 0},
		{"negative", -2, 3, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.Multiply(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	calc := Calculator{}

	tests := []struct {
		name        string
		a, b        int
		want        int
		expectError bool
	}{
		{"normal", 10, 2, 5, false},
		{"divide by zero", 10, 0, 0, true},
		{"negative", -10, 2, -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.Divide(tt.a, tt.b)

			if tt.expectError {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	calc := Calculator{}
	for i := 0; i < b.N; i++ {
		calc.Add(i, i+1)
	}
}

func FuzzDivide(f *testing.F) {
	calc := Calculator{}
	f.Add(10, 2)
	f.Add(0, 1)

	f.Fuzz(func(t *testing.T, a, b int) {
		if b == 0 {
			return // skip division by zero
		}
		result, err := calc.Divide(a, b)
		if err != nil {
			t.Fatalf("unexpected error for %d/%d: %v", a, b, err)
		}
		if result*b != a {
			t.Errorf("Divide(%d, %d) = %d; %d * %d != %d", a, b, result, result, b, a)
		}
	})
}
