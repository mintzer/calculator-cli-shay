package main

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
		{"zero and positive", 0, 5, 5},
		{"positive and zero", 5, 0, 5},
		{"both zero", 0, 0, 0},
		{"fractional values", 1.5, 2.3, 3.8},
		{"large numbers", 1e10, 2e10, 3e10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Fatalf("Add(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 10, 4, 6},
		{"negative numbers", -2, -3, 1},
		{"mixed signs", -2, 3, -5},
		{"zero minus positive", 0, 5, -5},
		{"positive minus zero", 5, 0, 5},
		{"both zero", 0, 0, 0},
		{"fractional values", 5.5, 2.3, 3.2},
		{"result is negative", 3, 7, -4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Fatalf("Subtract(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 15},
		{"negative numbers", -2, -3, 6},
		{"mixed signs", -2, 3, -6},
		{"zero times positive", 0, 5, 0},
		{"positive times zero", 5, 0, 0},
		{"both zero", 0, 0, 0},
		{"fractional values", 2.5, 4, 10},
		{"one", 7, 1, 7},
		{"negative one", 7, -1, -7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Fatalf("Multiply(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a, b      float64
		expected  float64
		expectErr bool
	}{
		{"positive numbers", 20, 4, 5, false},
		{"negative numbers", -6, -3, 2, false},
		{"mixed signs", -6, 3, -2, false},
		{"fractional result", 7, 2, 3.5, false},
		{"fractional inputs", 2.5, 0.5, 5, false},
		{"zero numerator", 0, 5, 0, false},
		{"one", 9, 1, 9, false},
		{"divide by zero", 20, 0, 0, true},
		{"zero divided by zero", 0, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("Divide(%v, %v) expected error, got nil", tt.a, tt.b)
				}
				if err.Error() != "cannot divide by zero" {
					t.Fatalf("Divide(%v, %v) error = %q, want %q", tt.a, tt.b, err.Error(), "cannot divide by zero")
				}
			} else {
				if err != nil {
					t.Fatalf("Divide(%v, %v) unexpected error: %v", tt.a, tt.b, err)
				}
				if math.Abs(got-tt.expected) > 1e-9 {
					t.Fatalf("Divide(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
				}
			}
		})
	}
}
