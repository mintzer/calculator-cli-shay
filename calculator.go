package main

import "errors"

// Add returns the sum of a and b.
func Add(a, b float64) float64 {
	return a + b
}

// Sub returns the difference of a and b (a - b).
func Sub(a, b float64) float64 {
	return a - b
}

// Mul returns the product of a and b.
func Mul(a, b float64) float64 {
	return a * b
}

// Div returns the quotient of a and b (a / b).
// It returns an error if b is zero.
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}
