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
		{"decimal values", 1.5, 2.5, 4},
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
		{"negative result", 3, 5, -2},
		{"negative numbers", -5, -3, -2},
		{"zeros", 0, 0, 0},
		{"same values", 7, 7, 0},
		{"decimal values", 5.5, 2.5, 3},
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
		{"positive numbers", 3, 7, 21},
		{"negative numbers", -3, -7, 21},
		{"mixed signs", -3, 7, -21},
		{"multiply by zero", 5, 0, 0},
		{"multiply by one", 5, 1, 5},
		{"decimal values", 2.5, 4, 10},
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
		{"exact division", 10, 2, 5, false},
		{"non-exact division", 10, 3, 10.0 / 3.0, false},
		{"negative numbers", -10, -2, 5, false},
		{"mixed signs", -10, 2, -5, false},
		{"zero numerator", 0, 5, 0, false},
		{"decimal values", 7.5, 2.5, 3, false},
		{"divide by zero", 10, 0, 0, true},
		{"zero divided by zero", 0, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Div(tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Div(%v, %v) expected error, got nil", tt.a, tt.b)
				}
				if err != ErrDivideByZero {
					t.Errorf("Div(%v, %v) error = %v, want ErrDivideByZero", tt.a, tt.b, err)
				}
				return
			}
			if err != nil {
				t.Errorf("Div(%v, %v) unexpected error: %v", tt.a, tt.b, err)
				return
			}
			if got != tt.want {
				t.Errorf("Div(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestFormatResult(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  string
	}{
		{"whole number", 8.0, "8.0"},
		{"non-terminating decimal", 10.0 / 3.0, "3.3333333333333335"},
		{"zero", 0.0, "0.0"},
		{"negative whole", -5.0, "-5.0"},
		{"decimal", 1.5, "1.5"},
		{"large whole number", 100.0, "100.0"},
		{"small decimal", 0.1, "0.1"},
		{"floating point artifact", math.Float64frombits(0x3FD3333333333334), "0.30000000000000004"},
		{"negative zero", math.Copysign(0, -1), "0.0"},
		{"twenty one", 21.0, "21.0"},
		{"large decimal", 1000000.5, "1000000.5"},
		{"very large whole", 1e16, "1e+16"},
		{"boundary whole 1e15", 1e15, "1000000000000000.0"},
		{"small scientific", 1e-5, "1e-05"},
		{"boundary small 1e-4", 1e-4, "0.0001"},
		{"large whole 1e20", 1e20, "1e+20"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatResult(tt.input)
			if got != tt.want {
				t.Errorf("formatResult(%v) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
