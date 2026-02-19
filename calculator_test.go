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
		{"negative numbers", -3, -7, -10},
		{"mixed signs", -5, 3, -2},
		{"with zero", 0, 5, 5},
		{"both zero", 0, 0, 0},
		{"fractional", 5.5, 3.5, 9},
		{"large numbers", 1e15, 1e15, 2e15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
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
		{"negative result", 3, 7, -4},
		{"negative numbers", -3, -7, 4},
		{"with zero", 5, 0, 5},
		{"both zero", 0, 0, 0},
		{"fractional", 5.5, 2.5, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sub(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Sub(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
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
		{"with zero", 0, 5, 0},
		{"negative numbers", -3, -4, 12},
		{"mixed signs", -3, 4, -12},
		{"both zero", 0, 0, 0},
		{"fractional", 2.5, 4, 10},
		{"one", 1, 99, 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Mul(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Mul(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
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
		{"fractional result", 7, 3, 7.0 / 3.0, false},
		{"negative numbers", -12, -4, 3, false},
		{"mixed signs", -12, 4, -3, false},
		{"zero numerator", 0, 5, 0, false},
		{"divide by zero", 5, 0, 0, true},
		{"zero divided by zero", 0, 0, 0, true},
		{"fractional", 5.5, 2.5, 2.2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Div(tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Div(%v, %v) expected error, got nil", tt.a, tt.b)
				} else if err.Error() != "Cannot divide by zero" {
					t.Errorf("Div(%v, %v) error = %q, want %q", tt.a, tt.b, err.Error(), "Cannot divide by zero")
				}
			} else {
				if err != nil {
					t.Errorf("Div(%v, %v) unexpected error: %v", tt.a, tt.b, err)
				}
				if math.Abs(got-tt.want) > 1e-9 {
					t.Errorf("Div(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
				}
			}
		})
	}
}
