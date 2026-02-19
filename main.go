package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const usageText = `usage: calc [-h] {add,sub,mul,div} a b

Simple CLI Calculator

positional arguments:
  {add,sub,mul,div}  Operation to perform
  a                  First number
  b                  Second number

options:
  -h, --help         show this help message and exit`

// formatResult formats a float64 to match Python's str(float) output behavior.
// Python always shows at least one decimal place for float values (e.g., 8.0)
// and uses scientific notation only for very large (>= 1e16) or very small (< 1e-4) values.
func formatResult(f float64) string {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return strconv.FormatFloat(f, 'g', -1, 64)
	}

	abs := math.Abs(f)
	var s string
	if f == 0 {
		s = "0.0"
	} else if abs >= 1e16 || abs < 1e-4 {
		// Python uses scientific notation in this range
		s = strconv.FormatFloat(f, 'g', -1, 64)
	} else {
		// Use fixed notation with shortest round-trip representation
		s = strconv.FormatFloat(f, 'f', -1, 64)
	}

	// Python always shows at least one decimal place for floats
	if !strings.Contains(s, ".") && !strings.Contains(s, "e") && !strings.Contains(s, "E") {
		s += ".0"
	}

	return s
}

// run is the testable core of the CLI. It processes the given args (without the
// program name), writes output to stdout/stderr, and returns an exit code.
func run(args []string, stdout, stderr io.Writer) int {
	// Handle help flags anywhere in args
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			fmt.Fprintln(stdout, usageText)
			return 0
		}
	}

	// Detect unknown flags (anything starting with - that isn't -h/--help)
	// Python's argparse identifies these before positional arg parsing
	var unknownFlags []string
	var positionalArgs []string
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") && arg != "-h" && arg != "--help" {
			// Check if it looks like a negative number (digits/dots after -)
			// by attempting to parse as float
			if _, err := strconv.ParseFloat(arg, 64); err != nil {
				unknownFlags = append(unknownFlags, arg)
			} else {
				positionalArgs = append(positionalArgs, arg)
			}
		} else {
			positionalArgs = append(positionalArgs, arg)
		}
	}

	// If there are unknown flags, report them
	if len(unknownFlags) > 0 {
		fmt.Fprintln(stderr, "usage: calc [-h] {add,sub,mul,div} a b")
		fmt.Fprintf(stderr, "calc: error: unrecognized arguments: %s\n", strings.Join(unknownFlags, " "))
		return 2
	}

	// Validate argument count
	if len(positionalArgs) < 3 {
		fmt.Fprintln(stderr, "usage: calc [-h] {add,sub,mul,div} a b")
		switch {
		case len(positionalArgs) == 0:
			fmt.Fprintln(stderr, "calc: error: the following arguments are required: operation, a, b")
		case len(positionalArgs) == 1:
			fmt.Fprintln(stderr, "calc: error: the following arguments are required: a, b")
		case len(positionalArgs) == 2:
			fmt.Fprintln(stderr, "calc: error: the following arguments are required: b")
		}
		return 2
	}

	if len(positionalArgs) > 3 {
		fmt.Fprintln(stderr, "usage: calc [-h] {add,sub,mul,div} a b")
		fmt.Fprintf(stderr, "calc: error: unrecognized arguments: %s\n", strings.Join(positionalArgs[3:], " "))
		return 2
	}

	operation := positionalArgs[0]
	aStr := positionalArgs[1]
	bStr := positionalArgs[2]

	// Validate operation
	validOps := []string{"add", "sub", "mul", "div"}
	isValid := false
	for _, op := range validOps {
		if operation == op {
			isValid = true
			break
		}
	}
	if !isValid {
		fmt.Fprintln(stderr, "usage: calc [-h] {add,sub,mul,div} a b")
		fmt.Fprintf(stderr, "calc: error: argument operation: invalid choice: '%s' (choose from add, sub, mul, div)\n", operation)
		return 2
	}

	// Parse numeric arguments
	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		fmt.Fprintln(stderr, "usage: calc [-h] {add,sub,mul,div} a b")
		fmt.Fprintf(stderr, "calc: error: argument a: invalid float value: '%s'\n", aStr)
		return 2
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		fmt.Fprintln(stderr, "usage: calc [-h] {add,sub,mul,div} a b")
		fmt.Fprintf(stderr, "calc: error: argument b: invalid float value: '%s'\n", bStr)
		return 2
	}

	// Execute operation
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

	fmt.Fprintln(stdout, formatResult(result))
	return 0
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
