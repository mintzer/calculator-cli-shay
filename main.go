package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// usage prints the CLI usage information to stderr.
func usage() {
	fmt.Fprintln(os.Stderr, "Usage: calculator <operation> <a> <b>")
	fmt.Fprintln(os.Stderr, "Operations: add, sub, mul, div")
}

// formatResult formats a float64 to match Python's str() output for floats.
// Integer-valued results include a trailing ".0" (e.g., 8 becomes "8.0").
func formatResult(v float64) string {
	s := strconv.FormatFloat(v, 'f', -1, 64)
	if !strings.Contains(s, ".") {
		s += ".0"
	}
	return s
}

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		usage()
		os.Exit(1)
	}

	operation := args[0]

	a, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid number: %s\n", args[1])
		os.Exit(1)
	}

	b, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid number: %s\n", args[2])
		os.Exit(1)
	}

	var result float64

	switch operation {
	case "add":
		result = add(a, b)
	case "sub":
		result = subtract(a, b)
	case "mul":
		result = multiply(a, b)
	case "div":
		result, err = divide(a, b)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown operation: %s\n", operation)
		usage()
		os.Exit(1)
	}

	fmt.Println(formatResult(result))
}
