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
		{"fractional", 1.5, 2.5, 4},
		{"large numbers", 1e15, 1e15, 2e15},
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
		{"fractional", 5.5, 2.5, 3},
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
		{"zero", 5, 0, 0},
		{"fractional", 3.5, 2, 7},
		{"identity", 42, 1, 42},
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
		{"positive numbers", 20, 4, 5, false},
		{"fractional result", 7, 3, 7.0 / 3.0, false},
		{"negative numbers", -12, -3, 4, false},
		{"mixed signs", -12, 3, -4, false},
		{"zero numerator", 0, 5, 0, false},
		{"divide by zero", 5, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Div(tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Div(%g, %g) expected error, got nil", tt.a, tt.b)
				}
				if err.Error() != "Cannot divide by zero" {
					t.Errorf("Div(%g, %g) error = %q, want %q", tt.a, tt.b, err.Error(), "Cannot divide by zero")
				}
				return
			}
			if err != nil {
				t.Errorf("Div(%g, %g) unexpected error: %v", tt.a, tt.b, err)
				return
			}
			if math.Abs(got-tt.want) > 1e-10 {
				t.Errorf("Div(%g, %g) = %g, want %g", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
