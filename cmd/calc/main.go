// Command calc is a simple CLI calculator.
package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"

	calculator "calculator-cli-shay"
)

// usage prints a short help message to stderr.
func usage() {
	fmt.Fprintln(os.Stderr, "Usage: calc <operation> <a> <b>")
	fmt.Fprintln(os.Stderr, "Operations: add, sub, mul, div")
}

// formatResult formats a float64 for display, matching Python's default
// float formatting: integral values are printed with one decimal place
// (e.g. "8.0"), while non-integral values show full precision without
// trailing zeros (e.g. "5.6").
func formatResult(v float64) string {
	if v == math.Trunc(v) && !math.IsInf(v, 0) && !math.IsNaN(v) {
		return strconv.FormatFloat(v, 'f', 1, 64)
	}
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func main() {
	if len(os.Args) != 4 {
		usage()
		os.Exit(2)
	}

	op := os.Args[1]
	a, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid number %q\n", os.Args[2])
		usage()
		os.Exit(2)
	}

	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid number %q\n", os.Args[3])
		usage()
		os.Exit(2)
	}

	result, err := calculator.Compute(op, a, b)
	if err != nil {
		if errors.Is(err, calculator.ErrDivisionByZero) {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
		// Invalid operation
		fmt.Fprintf(os.Stderr, "Error: %s %q\n", err, op)
		usage()
		os.Exit(2)
	}

	fmt.Println(formatResult(result))
}
