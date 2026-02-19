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
		{"negative numbers", -5, -3, -8},
		{"mixed signs", -5, 3, -2},
		{"zeros", 0, 0, 0},
		{"with zero", 5, 0, 5},
		{"decimal values", 1.5, 2.5, 4},
		{"large numbers", 1e10, 1e10, 2e10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 10, 4, 6},
		{"negative result", 3, 5, -2},
		{"negative numbers", -5, -3, -2},
		{"zeros", 0, 0, 0},
		{"subtract zero", 5, 0, 5},
		{"from zero", 0, 5, -5},
		{"decimal values", 5.5, 2.5, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sub(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Sub(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 6, 7, 42},
		{"with zero", 5, 0, 0},
		{"negative numbers", -3, -4, 12},
		{"mixed signs", -3, 4, -12},
		{"one", 5, 1, 5},
		{"decimal values", 2.5, 4, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Mul(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Mul(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		name      string
		a, b      float64
		expected  float64
		expectErr bool
	}{
		{"even division", 20, 4, 5, false},
		{"decimal result", 10, 3, 10.0 / 3.0, false},
		{"negative numbers", -10, -2, 5, false},
		{"mixed signs", -10, 2, -5, false},
		{"decimal values", 7.5, 2.5, 3, false},
		{"divide zero", 0, 5, 0, false},
		{"divide by zero", 10, 0, 0, true},
		{"zero divide by zero", 0, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Div(tt.a, tt.b)
			if tt.expectErr {
				if err == nil {
					t.Errorf("Div(%v, %v) expected error, got nil", tt.a, tt.b)
				} else if err.Error() != "Cannot divide by zero" {
					t.Errorf("Div(%v, %v) error = %q, want %q", tt.a, tt.b, err.Error(), "Cannot divide by zero")
				}
			} else {
				if err != nil {
					t.Errorf("Div(%v, %v) unexpected error: %v", tt.a, tt.b, err)
				}
				if math.Abs(result-tt.expected) > 1e-15 {
					t.Errorf("Div(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
				}
			}
		})
	}
}
