package calculator

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b   float64
		expect float64
	}{
		{5, 3, 8},
		{-1, 1, 0},
		{2.5, 0.5, 3},
		{0, 0, 0},
		{-3.5, -2.5, -6},
		{1e10, 1e10, 2e10},
	}
	for _, tt := range tests {
		got := Add(tt.a, tt.b)
		if got != tt.expect {
			t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expect)
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		a, b   float64
		expect float64
	}{
		{10, 3, 7},
		{-1, 1, -2},
		{2.5, 0.5, 2},
		{0, 0, 0},
		{-3, -5, 2},
		{1, 1, 0},
	}
	for _, tt := range tests {
		got := Sub(tt.a, tt.b)
		if got != tt.expect {
			t.Errorf("Sub(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expect)
		}
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		a, b   float64
		expect float64
	}{
		{2, 4, 8},
		{-2, 3, -6},
		{2.5, 2, 5},
		{0, 100, 0},
		{-3, -4, 12},
		{0.1, 0.2, 0.020000000000000004}, // IEEE 754 precision
	}
	for _, tt := range tests {
		got := Mul(tt.a, tt.b)
		if got != tt.expect {
			t.Errorf("Mul(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expect)
		}
	}
}

func TestDiv(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		tests := []struct {
			a, b   float64
			expect float64
		}{
			{20, 4, 5},
			{-6, 3, -2},
			{7, 2, 3.5},
			{0, 5, 0},
			{1, 3, 1.0 / 3.0},
		}
		for _, tt := range tests {
			got, err := Div(tt.a, tt.b)
			if err != nil {
				t.Errorf("Div(%v, %v) returned unexpected error: %v", tt.a, tt.b, err)
				continue
			}
			if got != tt.expect {
				t.Errorf("Div(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expect)
			}
		}
	})

	t.Run("division_by_zero", func(t *testing.T) {
		_, err := Div(10, 0)
		if err == nil {
			t.Fatal("Div(10, 0) expected error, got nil")
		}
		if !errors.Is(err, ErrDivisionByZero) {
			t.Errorf("Div(10, 0) error = %v, want ErrDivisionByZero", err)
		}
	})
}

func TestCompute(t *testing.T) {
	t.Run("valid_operations", func(t *testing.T) {
		tests := []struct {
			op     string
			a, b   float64
			expect float64
		}{
			{"add", 5, 3, 8},
			{"sub", 10, 3, 7},
			{"mul", 2, 4, 8},
			{"div", 20, 4, 5},
		}
		for _, tt := range tests {
			got, err := Compute(tt.op, tt.a, tt.b)
			if err != nil {
				t.Errorf("Compute(%q, %v, %v) returned unexpected error: %v", tt.op, tt.a, tt.b, err)
				continue
			}
			if got != tt.expect {
				t.Errorf("Compute(%q, %v, %v) = %v, want %v", tt.op, tt.a, tt.b, got, tt.expect)
			}
		}
	})

	t.Run("division_by_zero", func(t *testing.T) {
		_, err := Compute("div", 20, 0)
		if err == nil {
			t.Fatal("Compute(\"div\", 20, 0) expected error, got nil")
		}
		if !errors.Is(err, ErrDivisionByZero) {
			t.Errorf("Compute(\"div\", 20, 0) error = %v, want ErrDivisionByZero", err)
		}
	})

	t.Run("invalid_operation", func(t *testing.T) {
		_, err := Compute("foo", 1, 2)
		if err == nil {
			t.Fatal("Compute(\"foo\", 1, 2) expected error, got nil")
		}
		if !errors.Is(err, ErrInvalidOperation) {
			t.Errorf("Compute(\"foo\", 1, 2) error = %v, want ErrInvalidOperation", err)
		}
	})
}
