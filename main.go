package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const usageText = `Usage: calc-go <operation> <a> <b>

A simple CLI calculator.

Operations:
  add    Add two numbers
  sub    Subtract b from a
  mul    Multiply two numbers
  div    Divide a by b`

// formatFloat formats a float64 to match Python's default float formatting.
// Whole numbers get a trailing ".0", and fractional numbers use the shortest
// representation that round-trips (equivalent to Python's str()/repr()).
func formatFloat(f float64) string {
	s := strconv.FormatFloat(f, 'g', -1, 64)
	// If the string already contains a dot or is special (inf/NaN), return as-is.
	if strings.ContainsAny(s, ".eEnN") {
		return s
	}
	// Whole number: append ".0" to match Python's float output (e.g. "8.0").
	return s + ".0"
}

// run is the testable entry point for the CLI. It takes command-line arguments
// (without the program name), stdout and stderr writers, and returns an exit code.
func run(args []string, stdout, stderr io.Writer) int {
	// Handle help flags.
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			fmt.Fprintln(stdout, usageText)
			return 0
		}
	}

	// Validate argument count.
	if len(args) != 3 {
		fmt.Fprintln(stderr, usageText)
		if len(args) == 0 {
			fmt.Fprintln(stderr, "Error: the following arguments are required: operation, a, b")
		} else if len(args) == 1 {
			fmt.Fprintln(stderr, "Error: the following arguments are required: a, b")
		} else if len(args) == 2 {
			fmt.Fprintln(stderr, "Error: the following arguments are required: b")
		} else {
			fmt.Fprintln(stderr, "Error: too many arguments")
		}
		return 2
	}

	operation := args[0]
	aStr := args[1]
	bStr := args[2]

	// Validate operation.
	validOps := map[string]bool{"add": true, "sub": true, "mul": true, "div": true}
	if !validOps[operation] {
		fmt.Fprintln(stderr, usageText)
		fmt.Fprintf(stderr, "Error: invalid operation '%s' (choose from add, sub, mul, div)\n", operation)
		return 2
	}

	// Parse numeric arguments.
	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		fmt.Fprintf(stderr, "Error: invalid float value for a: '%s'\n", aStr)
		return 2
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		fmt.Fprintf(stderr, "Error: invalid float value for b: '%s'\n", bStr)
		return 2
	}

	// Dispatch operation.
	var result float64
	switch operation {
	case "add":
		result = Add(a, b)
	case "sub":
		result = Sub(a, b)
	case "mul":
		result = Mul(a, b)
	case "div":
		result, err = Div(a, b)
		if err != nil {
			fmt.Fprintf(stderr, "Error: %s\n", err.Error())
			return 1
		}
	}

	fmt.Fprintln(stdout, formatFloat(result))
	return 0
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
