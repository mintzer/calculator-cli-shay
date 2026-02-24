package main

import "errors"

// add returns the sum of a and b.
func add(a, b float64) float64 {
	return a + b
}

// subtract returns the difference of a and b (a minus b).
func subtract(a, b float64) float64 {
	return a - b
}

// multiply returns the product of a and b.
func multiply(a, b float64) float64 {
	return a * b
}

// divide returns the quotient of a divided by b.
// It returns an error if b is zero.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}
