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
// Python's float formatting rules:
//   - Whole numbers always show ".0" suffix (e.g. "8.0")
//   - Scientific notation is used when exponent >= 16 or exponent < -4
//   - Otherwise, decimal notation with minimum digits for unique representation
//   - Scientific notation uses lowercase "e+" / "e-" with 2-digit exponent minimum
func formatResult(f float64) string {
	// Handle negative zero
	if f == 0 {
		f = 0
	}

	if math.IsInf(f, 0) || math.IsNaN(f) {
		return strconv.FormatFloat(f, 'g', -1, 64)
	}

	// Get the shortest unique representation using Go's 'e' and 'f' formats
	// then choose based on Python's exponent thresholds.
	// First, determine the exponent by using the 'e' format.
	eStr := strconv.FormatFloat(f, 'e', -1, 64)

	// Parse exponent from the 'e' representation (e.g., "1.0000005e+06")
	exp := 0
	expNeg := false
	eIdx := len(eStr) - 1
	for eIdx >= 0 && eStr[eIdx] != 'e' {
		eIdx--
	}
	if eIdx >= 0 {
		expPart := eStr[eIdx+1:]
		if len(expPart) > 0 && expPart[0] == '+' {
			expPart = expPart[1:]
		} else if len(expPart) > 0 && expPart[0] == '-' {
			expNeg = true
			expPart = expPart[1:]
		}
		for _, c := range expPart {
			exp = exp*10 + int(c-'0')
		}
		if expNeg {
			exp = -exp
		}
	}

	// Python uses scientific notation for exponent >= 16 or < -4
	if exp >= 16 || exp < -4 {
		// Format in scientific notation matching Python's style
		// Python uses e.g. "1e+16", "1e-05" (2-digit minimum exponent)
		s := strconv.FormatFloat(f, 'e', -1, 64)
		return formatPythonSciNotation(s)
	}

	// For values in the decimal range, use 'f' format with minimum digits
	// If it's a whole number, show exactly one decimal place
	if f == math.Trunc(f) {
		return strconv.FormatFloat(f, 'f', 1, 64)
	}

	// For non-whole numbers, use 'f' format with enough precision
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// formatPythonSciNotation converts Go's scientific notation (e.g. "1e+06")
// to Python's format (e.g. "1e+06"). Go already uses the same format,
// but we ensure the exponent has at least 2 digits with a leading zero.
func formatPythonSciNotation(s string) string {
	// Find the 'e' character
	eIdx := -1
	for i, c := range s {
		if c == 'e' {
			eIdx = i
			break
		}
	}
	if eIdx < 0 {
		return s
	}

	mantissa := s[:eIdx]
	expPart := s[eIdx+1:] // e.g. "+06" or "-05" or "+6"

	sign := "+"
	digits := expPart
	if len(expPart) > 0 && (expPart[0] == '+' || expPart[0] == '-') {
		sign = string(expPart[0])
		digits = expPart[1:]
	}

	// Remove leading zeros then re-pad to minimum 2 digits
	i := 0
	for i < len(digits)-1 && digits[i] == '0' {
		i++
	}
	digits = digits[i:]
	if len(digits) < 2 {
		digits = "0" + digits
	}

	return mantissa + "e" + sign + digits
}
