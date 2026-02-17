package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const usageLine = "usage: calc [-h] {add,sub,mul,div} a b"

const fullHelp = `usage: calc [-h] {add,sub,mul,div} a b

Simple CLI Calculator

positional arguments:
  {add,sub,mul,div}  Operation to perform
  a                  First number
  b                  Second number

options:
  -h, --help         show this help message and exit
`

var validOps = []string{"add", "sub", "mul", "div"}

func isValidOp(op string) bool {
	for _, v := range validOps {
		if op == v {
			return true
		}
	}
	return false
}

// formatFloat formats a float64 the same way Python's default float formatting does.
// Python prints e.g. 8.0, -5.0, 3.5, 0.003, -0.0, 1000000000.0, 105.0
func formatFloat(f float64) string {
	// Python's default str(float) / f"{result}" behavior:
	// - integers get .0 suffix (8.0, -5.0)
	// - fractional numbers shown normally (3.5, 0.003)
	// - very large numbers may use scientific notation but in our test range they don't
	// - negative zero is -0.0

	// Use Go's default %g-like formatting but we need to match Python exactly.
	// Python uses repr-style: str(8.0) = "8.0", str(3.5) = "3.5", str(0.003) = "0.003"
	// str(1000000000.0) = "1000000000.0", str(-0.0) = "-0.0"

	// Python uses %g-like but with higher precision and always shows at least one decimal for floats.
	// Actually Python's default float.__str__ uses David Gay's algorithm (like Go's strconv).
	// The key difference: Python always includes a decimal point for float output.

	s := strconv.FormatFloat(f, 'f', -1, 64)

	// If there's no decimal point, add .0 (matching Python behavior)
	if !strings.Contains(s, ".") {
		s += ".0"
	}

	return s
}

func argparseError(msg string) {
	fmt.Fprintf(os.Stderr, "%s\ncalc: error: %s\n", usageLine, msg)
	os.Exit(2)
}

func main() {
	args := os.Args[1:]

	// Check for help flag anywhere in args
	for _, a := range args {
		if a == "-h" || a == "--help" {
			fmt.Print(fullHelp)
			os.Exit(0)
		}
	}

	// Filter out unknown flags and collect positional args, mimicking argparse behavior.
	// Python's argparse with positional-only args treats unknown flags as unrecognized.
	// But the actual SRC behavior for --unknown and -z with no positional args is:
	// "the following arguments are required: operation, a, b" (exit 2)
	// This means argparse checks required positionals first before complaining about unknown flags.

	// Separate positional args from flag-like args
	var positional []string
	var unrecognized []string
	for _, a := range args {
		if strings.HasPrefix(a, "-") && a != "-" {
			// Could be a negative number if it parses as float and we're past the operation
			// But argparse treats this differently: the operation is a choices field,
			// and a/b are float types. Argparse processes left-to-right.
			// Actually for argparse, if we haven't consumed the operation yet, -z would be a flag.
			// But with choices for operation, -z is not a valid choice either.
			// The SRC behavior shows: --unknown with no other args => "arguments are required: operation, a, b"
			// This means argparse prioritizes missing required args over unrecognized flags.

			// However, negative numbers like "-2" need to work as operands.
			// In argparse, negative numbers work because argparse recognizes them.
			// We need to figure out context: if this looks like a negative number and we're
			// expecting a number (position >= 1), treat as positional.
			if len(positional) >= 1 {
				// We're in operand position, try to parse as float
				if _, err := strconv.ParseFloat(a, 64); err == nil {
					positional = append(positional, a)
					continue
				}
			}
			unrecognized = append(unrecognized, a)
		} else {
			positional = append(positional, a)
		}
	}

	// Determine missing required args (mimicking argparse order)
	// argparse checks required positionals before reporting unrecognized args
	var missing []string
	if len(positional) < 1 {
		missing = append(missing, "operation")
	}
	if len(positional) < 2 {
		missing = append(missing, "a")
	}
	if len(positional) < 3 {
		missing = append(missing, "b")
	}

	if len(missing) > 0 {
		argparseError(fmt.Sprintf("the following arguments are required: %s", strings.Join(missing, ", ")))
	}

	// Check for too many positional args
	if len(positional) > 3 {
		extra := positional[3:]
		argparseError(fmt.Sprintf("unrecognized arguments: %s", strings.Join(extra, " ")))
	}

	// Check for unrecognized flag args
	if len(unrecognized) > 0 {
		argparseError(fmt.Sprintf("unrecognized arguments: %s", strings.Join(unrecognized, " ")))
	}

	operation := positional[0]
	aStr := positional[1]
	bStr := positional[2]

	// Validate operation (choices)
	if !isValidOp(operation) {
		argparseError(fmt.Sprintf("argument operation: invalid choice: '%s' (choose from add, sub, mul, div)", operation))
	}

	// Parse first number
	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		argparseError(fmt.Sprintf("argument a: invalid float value: '%s'", aStr))
	}

	// Parse second number
	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		argparseError(fmt.Sprintf("argument b: invalid float value: '%s'", bStr))
	}

	// Execute operation
	var result float64
	switch operation {
	case "add":
		result = Add(a, b)
	case "sub":
		result = Subtract(a, b)
	case "mul":
		result = Multiply(a, b)
	case "div":
		val, divErr := Divide(a, b)
		if divErr != nil {
			fmt.Fprintf(os.Stderr, "Error: Cannot divide by zero\n")
			os.Exit(1)
		}
		result = val
	}

	// Handle negative zero: math.Signbit detects -0.0
	// Python outputs "-0.0" for negative zero
	if result == 0 && math.Signbit(result) {
		fmt.Println("-0.0")
	} else {
		fmt.Println(formatFloat(result))
	}
}
