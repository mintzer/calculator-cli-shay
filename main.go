package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const usageText = `Usage: calc-go <operation> <a> <b>

Simple CLI Calculator

Positional arguments:
  operation    Operation to perform: add, sub, mul, div
  a            First number
  b            Second number`

// formatResult formats a float64 to match Python's default str(float) behavior.
func formatResult(f float64) string {
	// Use 'f' format for whole numbers to avoid Go's early scientific notation
	if f == math.Trunc(f) && !math.IsInf(f, 0) && !math.IsNaN(f) {
		return strconv.FormatFloat(f, 'f', 1, 64)
	}
	// For non-whole numbers, use 'g' with -1 precision (shortest unique representation)
	s := strconv.FormatFloat(f, 'g', -1, 64)
	if !strings.Contains(s, ".") && !strings.Contains(s, "e") && !strings.Contains(s, "E") {
		return s + ".0"
	}
	return s
}

// run executes the CLI logic with the given arguments and I/O writers.
// Returns the exit code.
func run(args []string, stdout, stderr io.Writer) int {
	// Check for help flags or no arguments
	if len(args) == 0 {
		fmt.Fprintln(stderr, usageText)
		return 2
	}

	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			fmt.Fprintln(stdout, usageText)
			return 0
		}
	}

	if len(args) < 3 {
		fmt.Fprintln(stderr, "Error: requires exactly 3 arguments: <operation> <a> <b>")
		fmt.Fprintln(stderr, usageText)
		return 2
	}

	if len(args) > 3 {
		fmt.Fprintln(stderr, "Error: too many arguments, requires exactly 3: <operation> <a> <b>")
		fmt.Fprintln(stderr, usageText)
		return 2
	}

	operation := args[0]
	aStr := args[1]
	bStr := args[2]

	// Validate operation
	validOps := map[string]bool{"add": true, "sub": true, "mul": true, "div": true}
	if !validOps[operation] {
		fmt.Fprintf(stderr, "Error: invalid operation '%s' (choose from add, sub, mul, div)\n", operation)
		return 2
	}

	// Parse numeric arguments
	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		fmt.Fprintf(stderr, "Error: invalid number for a: '%s'\n", aStr)
		return 2
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		fmt.Fprintf(stderr, "Error: invalid number for b: '%s'\n", bStr)
		return 2
	}

	// Dispatch to operation
	var result float64
	var opErr error

	switch operation {
	case "add":
		result = Add(a, b)
	case "sub":
		result = Sub(a, b)
	case "mul":
		result = Mul(a, b)
	case "div":
		result, opErr = Div(a, b)
	}

	if opErr != nil {
		fmt.Fprintf(stderr, "Error: %s\n", opErr.Error())
		return 1
	}

	fmt.Fprintln(stdout, formatResult(result))
	return 0
}

func main() {
	exitCode := run(os.Args[1:], os.Stdout, os.Stderr)
	os.Exit(exitCode)
}
