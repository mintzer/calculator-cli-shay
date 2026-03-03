// Command calc is a simple CLI calculator.
package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	calculator "calculator-cli-shay"
)

const usageLine = "usage: calc [-h] {add,sub,mul,div} a b"

const helpText = `usage: calc [-h] {add,sub,mul,div} a b

Simple CLI Calculator

positional arguments:
  {add,sub,mul,div}  Operation to perform
  a                  First number
  b                  Second number

options:
  -h, --help         show this help message and exit
`

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

// looksLikeFlag returns true if arg starts with '-' and is not a negative number.
func looksLikeFlag(arg string) bool {
	if !strings.HasPrefix(arg, "-") {
		return false
	}
	_, err := strconv.ParseFloat(arg, 64)
	return err != nil
}

func exitError(msg string) {
	fmt.Fprintln(os.Stderr, usageLine)
	fmt.Fprintf(os.Stderr, "calc: error: %s\n", msg)
	os.Exit(2)
}

func main() {
	args := os.Args[1:]

	// Check for help flags first (argparse handles -h/--help before anything else)
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			fmt.Print(helpText)
			os.Exit(0)
		}
	}

	// Separate flags from positional args
	var flags []string
	var positional []string
	for _, arg := range args {
		if looksLikeFlag(arg) {
			flags = append(flags, arg)
		} else {
			positional = append(positional, arg)
		}
	}

	// Check for missing required positional args (argparse checks this before unrecognized)
	if len(positional) < 3 {
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
		exitError(fmt.Sprintf("the following arguments are required: %s", strings.Join(missing, ", ")))
	}

	// Check for unrecognized flags
	if len(flags) > 0 {
		exitError(fmt.Sprintf("unrecognized arguments: %s", strings.Join(flags, " ")))
	}

	// Check for extra positional args
	if len(positional) > 3 {
		extra := make([]string, len(positional)-3)
		copy(extra, positional[3:])
		exitError(fmt.Sprintf("unrecognized arguments: %s", strings.Join(extra, " ")))
	}

	op := positional[0]
	aStr := positional[1]
	bStr := positional[2]

	// Validate operation choice
	validOps := map[string]bool{"add": true, "sub": true, "mul": true, "div": true}
	if !validOps[op] {
		exitError(fmt.Sprintf("argument operation: invalid choice: '%s' (choose from add, sub, mul, div)", op))
	}

	// Parse operands
	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		exitError(fmt.Sprintf("argument a: invalid float value: '%s'", aStr))
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		exitError(fmt.Sprintf("argument b: invalid float value: '%s'", bStr))
	}

	// Compute
	result, err := calculator.Compute(op, a, b)
	if err != nil {
		if errors.Is(err, calculator.ErrDivisionByZero) {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
		exitError(err.Error())
	}

	fmt.Println(formatResult(result))
}
