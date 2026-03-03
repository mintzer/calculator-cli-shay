// Package calculator provides basic arithmetic operations.
package calculator

import "errors"

// Sentinel errors for error identification.
var (
	ErrDivisionByZero  = errors.New("Cannot divide by zero")
	ErrInvalidOperation = errors.New("invalid operation")
)

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

// Div returns the quotient a / b. It returns ErrDivisionByZero if b is zero.
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// Compute dispatches the operation op on operands a and b.
// Supported operations: "add", "sub", "mul", "div".
// Returns ErrInvalidOperation for unrecognized operation names.
func Compute(op string, a, b float64) (float64, error) {
	switch op {
	case "add":
		return Add(a, b), nil
	case "sub":
		return Sub(a, b), nil
	case "mul":
		return Mul(a, b), nil
	case "div":
		return Div(a, b)
	default:
		return 0, ErrInvalidOperation
	}
}
