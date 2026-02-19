package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const progName = "calc"

const usageLine = "usage: " + progName + " [-h] {add,sub,mul,div} a b"

const helpText = usageLine + `

Simple CLI Calculator

positional arguments:
  {add,sub,mul,div}  Operation to perform
  a                  First number
  b                  Second number

options:
  -h, --help         show this help message and exit`

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
	// Separate flags from positional arguments
	var positional []string
	var unrecognized []string
	helpRequested := false

	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			helpRequested = true
		} else if strings.HasPrefix(arg, "-") {
			// Check if it looks like a negative number
			if len(arg) > 1 {
				_, err := strconv.ParseFloat(arg, 64)
				if err == nil {
					positional = append(positional, arg)
					continue
				}
			}
			unrecognized = append(unrecognized, arg)
		} else {
			positional = append(positional, arg)
		}
	}

	// Handle help
	if helpRequested {
		fmt.Fprintln(stdout, helpText)
		return 0
	}

	// Check required positional arguments first (matches Python argparse priority)
	validOps := map[string]bool{"add": true, "sub": true, "mul": true, "div": true}

	if len(positional) == 0 {
		fmt.Fprintln(stderr, usageLine)
		fmt.Fprintf(stderr, "%s: error: the following arguments are required: operation, a, b\n", progName)
		return 2
	}

	// Handle unrecognized arguments (flags that aren't -h/--help)
	if len(unrecognized) > 0 {
		fmt.Fprintln(stderr, usageLine)
		fmt.Fprintf(stderr, "%s: error: unrecognized arguments: %s\n", progName, strings.Join(unrecognized, " "))
		return 2
	}

	operation := positional[0]

	// Validate operation choice
	if !validOps[operation] {
		fmt.Fprintln(stderr, usageLine)
		fmt.Fprintf(stderr, "%s: error: argument operation: invalid choice: '%s' (choose from add, sub, mul, div)\n", progName, operation)
		return 2
	}

	if len(positional) == 1 {
		fmt.Fprintln(stderr, usageLine)
		fmt.Fprintf(stderr, "%s: error: the following arguments are required: a, b\n", progName)
		return 2
	}

	if len(positional) == 2 {
		fmt.Fprintln(stderr, usageLine)
		fmt.Fprintf(stderr, "%s: error: the following arguments are required: b\n", progName)
		return 2
	}

	if len(positional) > 3 {
		extra := strings.Join(positional[3:], " ")
		fmt.Fprintln(stderr, usageLine)
		fmt.Fprintf(stderr, "%s: error: unrecognized arguments: %s\n", progName, extra)
		return 2
	}

	aStr := positional[1]
	bStr := positional[2]

	// Parse numeric arguments (matching Python argparse float validation)
	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		fmt.Fprintln(stderr, usageLine)
		fmt.Fprintf(stderr, "%s: error: argument a: invalid float value: '%s'\n", progName, aStr)
		return 2
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		fmt.Fprintln(stderr, usageLine)
		fmt.Fprintf(stderr, "%s: error: argument b: invalid float value: '%s'\n", progName, bStr)
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
