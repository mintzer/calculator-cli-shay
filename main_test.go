package main

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
		{"decimals", 1.5, 2.3, 3.8},
		{"zero", 0, 5, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := add(tt.a, tt.b)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("add(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 5, 3, 2},
		{"negative numbers", -2, -3, 1},
		{"mixed signs", -2, 3, -5},
		{"decimals", 5.5, 2.3, 3.2},
		{"zero", 5, 0, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subtract(tt.a, tt.b)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("subtract(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 5, 3, 15},
		{"negative numbers", -2, -3, 6},
		{"mixed signs", -2, 3, -6},
		{"decimals", 1.5, 2.0, 3.0},
		{"multiply by zero", 5, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := multiply(tt.a, tt.b)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("multiply(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{"positive numbers", 6, 3, 2, false},
		{"negative numbers", -6, -3, 2, false},
		{"mixed signs", -6, 3, -2, false},
		{"decimals", 7.5, 2.5, 3.0, false},
		{"non-integer result", 10, 3, 10.0 / 3.0, false},
		{"divide by zero", 5, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide(tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Errorf("divide(%v, %v) expected error, got nil", tt.a, tt.b)
				}
				if err.Error() != "Cannot divide by zero" {
					t.Errorf("divide(%v, %v) error = %q, want %q", tt.a, tt.b, err.Error(), "Cannot divide by zero")
				}
			} else {
				if err != nil {
					t.Errorf("divide(%v, %v) unexpected error: %v", tt.a, tt.b, err)
				}
				if math.Abs(got-tt.want) > 1e-9 {
					t.Errorf("divide(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
				}
			}
		})
	}
}

func TestDivideByZeroNoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("divide(1, 0) panicked: %v", r)
		}
	}()
	_, err := divide(1, 0)
	if err == nil {
		t.Error("divide(1, 0) expected error, got nil")
	}
}
