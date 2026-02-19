package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const usage = `Usage: calc-go <operation> <a> <b>

Simple CLI Calculator

Operations:
  add    Add two numbers
  sub    Subtract b from a
  mul    Multiply two numbers
  div    Divide a by b`

// formatFloat formats a float64 to match Python's default float string
// representation: minimum digits to uniquely identify the value, with at
// least one decimal place always shown (e.g., 8.0, not 8).
func formatFloat(v float64) string {
	s := strconv.FormatFloat(v, 'g', -1, 64)
	if !strings.Contains(s, ".") && !strings.Contains(s, "e") && !strings.Contains(s, "E") {
		s += ".0"
	}
	return s
}

// run executes the CLI logic with the given arguments and I/O writers.
// It returns the exit code.
func run(args []string, stdout, stderr io.Writer) int {
	// Handle help flags
	if len(args) == 1 && (args[0] == "-h" || args[0] == "--help") {
		fmt.Fprintln(stdout, usage)
		return 0
	}

	// Validate argument count
	if len(args) != 3 {
		fmt.Fprintln(stderr, usage)
		if len(args) == 0 {
			fmt.Fprintln(stderr, "Error: the following arguments are required: operation, a, b")
		} else {
			fmt.Fprintln(stderr, "Error: expected exactly 3 arguments: operation, a, b")
		}
		return 2
	}

	operation := args[0]
	aStr := args[1]
	bStr := args[2]

	// Validate operation
	validOps := map[string]bool{"add": true, "sub": true, "mul": true, "div": true}
	if !validOps[operation] {
		fmt.Fprintln(stderr, usage)
		fmt.Fprintf(stderr, "Error: invalid operation '%s' (choose from add, sub, mul, div)\n", operation)
		return 2
	}

	// Parse operands
	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		fmt.Fprintln(stderr, usage)
		fmt.Fprintf(stderr, "Error: invalid float value for a: '%s'\n", aStr)
		return 2
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		fmt.Fprintln(stderr, usage)
		fmt.Fprintf(stderr, "Error: invalid float value for b: '%s'\n", bStr)
		return 2
	}

	// Dispatch operation
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
	exitCode := run(os.Args[1:], os.Stdout, os.Stderr)
	os.Exit(exitCode)
}
