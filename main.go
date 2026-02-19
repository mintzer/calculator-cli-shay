package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const usageLine = "usage: calc [-h] {add,sub,mul,div} a b"

const helpText = `usage: calc [-h] {add,sub,mul,div} a b

Simple CLI Calculator

positional arguments:
  {add,sub,mul,div}  Operation to perform
  a                  First number
  b                  Second number

options:
  -h, --help         show this help message and exit`

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

// cliError prints the usage line followed by "calc: error: <msg>" to stderr
// and returns exit code 2, matching Python argparse error output.
func cliError(stderr io.Writer, msg string) int {
	fmt.Fprintln(stderr, usageLine)
	fmt.Fprintf(stderr, "calc: error: %s\n", msg)
	return 2
}

// run is the testable entry point for the CLI. It takes command-line arguments
// (without the program name), stdout and stderr writers, and returns an exit code.
func run(args []string, stdout, stderr io.Writer) int {
	// Handle help flags.
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			fmt.Fprintln(stdout, helpText)
			return 0
		}
	}

	// Filter out unknown flags/options (anything starting with - that isn't a number),
	// matching Python argparse behavior which ignores unknown options but treats
	// negative numbers as positional arguments.
	var positional []string
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") && arg != "-" {
			// Check if it's a negative number (starts with - followed by digit or dot)
			rest := arg[1:]
			if len(rest) > 0 && (rest[0] >= '0' && rest[0] <= '9' || rest[0] == '.') {
				positional = append(positional, arg)
				continue
			}
			// Unknown flag — ignore it (Python argparse skips these for positional parsing)
			continue
		}
		positional = append(positional, arg)
	}

	validOps := map[string]bool{"add": true, "sub": true, "mul": true, "div": true}

	// Validate we have enough positional arguments.
	if len(positional) < 3 {
		var missing []string
		if len(positional) == 0 {
			missing = append(missing, "operation", "a", "b")
		} else if len(positional) == 1 {
			// Check if the operation is valid first
			if !validOps[positional[0]] {
				return cliError(stderr, fmt.Sprintf("argument operation: invalid choice: '%s' (choose from add, sub, mul, div)", positional[0]))
			}
			missing = append(missing, "a", "b")
		} else if len(positional) == 2 {
			if !validOps[positional[0]] {
				return cliError(stderr, fmt.Sprintf("argument operation: invalid choice: '%s' (choose from add, sub, mul, div)", positional[0]))
			}
			missing = append(missing, "b")
		}
		return cliError(stderr, fmt.Sprintf("the following arguments are required: %s", strings.Join(missing, ", ")))
	}

	// Check for extra arguments beyond the three expected.
	if len(positional) > 3 {
		return cliError(stderr, fmt.Sprintf("unrecognized arguments: %s", strings.Join(positional[3:], " ")))
	}

	operation := positional[0]
	aStr := positional[1]
	bStr := positional[2]

	// Validate operation.
	if !validOps[operation] {
		return cliError(stderr, fmt.Sprintf("argument operation: invalid choice: '%s' (choose from add, sub, mul, div)", operation))
	}

	// Parse numeric arguments.
	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		return cliError(stderr, fmt.Sprintf("argument a: invalid float value: '%s'", aStr))
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		return cliError(stderr, fmt.Sprintf("argument b: invalid float value: '%s'", bStr))
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
