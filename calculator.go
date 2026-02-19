package main

import (
	"errors"
	"math"
	"strconv"
)

// ErrDivideByZero is returned when division by zero is attempted.
var ErrDivideByZero = errors.New("Cannot divide by zero")

// Add returns the sum of a and b.
func Add(a, b float64) float64 {
	return a + b
}

// Sub returns the difference a - b.
func Sub(a, b float64) float64 {
	return a - b
}

// Mul returns the product of a and b.
func Mul(a, b float64) float64 {
	return a * b
}

// Div returns a / b. Returns an error if b is zero.
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// formatResult formats a float64 to match Python's str(float) output.
// Python always shows at least one decimal place for floats (e.g. "8.0")
// and uses up to 17 significant digits for non-whole numbers.
func formatResult(f float64) string {
	// Handle negative zero
	if f == 0 {
		f = 0
	}

	// If the value is a whole number (no fractional part), format with .0
	if f == math.Trunc(f) && !math.IsInf(f, 0) && !math.IsNaN(f) {
		return strconv.FormatFloat(f, 'f', 1, 64)
	}

	// For non-whole numbers, use Go's default representation which matches
	// Python's repr-style output for float64 values.
	// strconv.FormatFloat with 'g' and -1 precision uses the minimum number
	// of digits necessary to represent the value uniquely, matching Python's behavior.
	return strconv.FormatFloat(f, 'g', -1, 64)
}
