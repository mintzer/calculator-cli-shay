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
		{"negative numbers", -5, -3, -8},
		{"mixed signs", -5, 3, -2},
		{"zeros", 0, 0, 0},
		{"float values", 1.5, 2.5, 4},
		{"large numbers", 1e10, 2e10, 3e10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%g, %g) = %g, want %g", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 10, 4, 6},
		{"negative result", 3, 5, -2},
		{"negative numbers", -5, -3, -2},
		{"zeros", 0, 0, 0},
		{"float values", 5.5, 2.5, 3},
		{"same numbers", 7, 7, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sub(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Sub(%g, %g) = %g, want %g", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 6, 7, 42},
		{"negative numbers", -3, -4, 12},
		{"mixed signs", -3, 4, -12},
		{"multiply by zero", 5, 0, 0},
		{"multiply by one", 5, 1, 5},
		{"float values", 2.5, 4, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Mul(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Mul(%g, %g) = %g, want %g", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{"exact division", 20, 4, 5, false},
		{"non-exact division", 10, 3, 10.0 / 3.0, false},
		{"negative numbers", -12, -4, 3, false},
		{"mixed signs", -12, 4, -3, false},
		{"divide zero", 0, 5, 0, false},
		{"float values", 7.5, 2.5, 3, false},
		{"divide by zero", 10, 0, 0, true},
		{"divide zero by zero", 0, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Div(tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Div(%g, %g) expected error, got nil", tt.a, tt.b)
				}
				if err != nil && err.Error() != "Cannot divide by zero" {
					t.Errorf("Div(%g, %g) error = %q, want %q", tt.a, tt.b, err.Error(), "Cannot divide by zero")
				}
			} else {
				if err != nil {
					t.Errorf("Div(%g, %g) unexpected error: %v", tt.a, tt.b, err)
				}
				if math.Abs(got-tt.want) > 1e-10 {
					t.Errorf("Div(%g, %g) = %g, want %g", tt.a, tt.b, got, tt.want)
				}
			}
		})
	}
}
