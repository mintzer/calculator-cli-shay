package main

import (
	"math"
	"testing"
)

// almostEqual checks whether two float64 values are within epsilon of each other.
func almostEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

// --- Unit tests for arithmetic functions ---

func TestAddBasic(t *testing.T) {
	got := add(5, 3)
	want := 8.0
	if got != want {
		t.Fatalf("add(5, 3) = %v, want %v", got, want)
	}
}

func TestAddNegative(t *testing.T) {
	got := add(-5, 3)
	want := -2.0
	if got != want {
		t.Fatalf("add(-5, 3) = %v, want %v", got, want)
	}
}

func TestAddFloatingPoint(t *testing.T) {
	got := add(0.1, 0.2)
	want := 0.3
	if !almostEqual(got, want, 1e-9) {
		t.Fatalf("add(0.1, 0.2) = %v, want approximately %v", got, want)
	}
}

func TestAddZeros(t *testing.T) {
	got := add(0, 0)
	want := 0.0
	if got != want {
		t.Fatalf("add(0, 0) = %v, want %v", got, want)
	}
}

func TestSubtractBasic(t *testing.T) {
	got := subtract(10, 4)
	want := 6.0
	if got != want {
		t.Fatalf("subtract(10, 4) = %v, want %v", got, want)
	}
}

func TestSubtractNegative(t *testing.T) {
	got := subtract(-10, -4)
	want := -6.0
	if got != want {
		t.Fatalf("subtract(-10, -4) = %v, want %v", got, want)
	}
}

func TestSubtractResultNegative(t *testing.T) {
	got := subtract(3, 10)
	want := -7.0
	if got != want {
		t.Fatalf("subtract(3, 10) = %v, want %v", got, want)
	}
}

func TestMultiplyBasic(t *testing.T) {
	got := multiply(6, 7)
	want := 42.0
	if got != want {
		t.Fatalf("multiply(6, 7) = %v, want %v", got, want)
	}
}

func TestMultiplyNegative(t *testing.T) {
	got := multiply(-6, 7)
	want := -42.0
	if got != want {
		t.Fatalf("multiply(-6, 7) = %v, want %v", got, want)
	}
}

func TestMultiplyByZero(t *testing.T) {
	got := multiply(100, 0)
	want := 0.0
	if got != want {
		t.Fatalf("multiply(100, 0) = %v, want %v", got, want)
	}
}

func TestDivideBasic(t *testing.T) {
	got, err := divide(20, 4)
	if err != nil {
		t.Fatalf("divide(20, 4) returned unexpected error: %v", err)
	}
	want := 5.0
	if got != want {
		t.Fatalf("divide(20, 4) = %v, want %v", got, want)
	}
}

func TestDivideNegative(t *testing.T) {
	got, err := divide(-20, 4)
	if err != nil {
		t.Fatalf("divide(-20, 4) returned unexpected error: %v", err)
	}
	want := -5.0
	if got != want {
		t.Fatalf("divide(-20, 4) = %v, want %v", got, want)
	}
}

func TestDivideFractionalResult(t *testing.T) {
	got, err := divide(7, 2)
	if err != nil {
		t.Fatalf("divide(7, 2) returned unexpected error: %v", err)
	}
	want := 3.5
	if got != want {
		t.Fatalf("divide(7, 2) = %v, want %v", got, want)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := divide(1, 0)
	if err == nil {
		t.Fatal("divide(1, 0) expected an error, got nil")
	}
	expectedMsg := "Cannot divide by zero"
	if err.Error() != expectedMsg {
		t.Fatalf("divide(1, 0) error = %q, want %q", err.Error(), expectedMsg)
	}
}

// --- Unit tests for formatResult ---

func TestFormatResultInteger(t *testing.T) {
	got := formatResult(8.0)
	want := "8.0"
	if got != want {
		t.Fatalf("formatResult(8.0) = %q, want %q", got, want)
	}
}

func TestFormatResultNegativeInteger(t *testing.T) {
	got := formatResult(-2.0)
	want := "-2.0"
	if got != want {
		t.Fatalf("formatResult(-2.0) = %q, want %q", got, want)
	}
}

func TestFormatResultFractional(t *testing.T) {
	got := formatResult(5.5)
	want := "5.5"
	if got != want {
		t.Fatalf("formatResult(5.5) = %q, want %q", got, want)
	}
}

func TestFormatResultZero(t *testing.T) {
	got := formatResult(0.0)
	want := "0.0"
	if got != want {
		t.Fatalf("formatResult(0.0) = %q, want %q", got, want)
	}
}
