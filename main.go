package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

const usageText = `Usage: calc-go <operation> <a> <b>

A simple CLI calculator.

Operations:
  add    Add two numbers
  sub    Subtract b from a
  mul    Multiply two numbers
  div    Divide a by b
`

// run executes the calculator CLI logic. It takes command-line arguments
// (without the program name), stdout and stderr writers, and returns an
// exit code. This design enables easy testing without process spawning.
func run(args []string, stdout, stderr io.Writer) int {
	// Handle -h / --help flags
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			fmt.Fprint(stdout, usageText)
			return 0
		}
	}

	// Validate argument count
	if len(args) != 3 {
		fmt.Fprint(stderr, usageText)
		fmt.Fprintf(stderr, "Error: expected 3 arguments, got %d\n", len(args))
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

	// Parse operands
	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		fmt.Fprintf(stderr, "Error: invalid number '%s' for argument a\n", aStr)
		return 2
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		fmt.Fprintf(stderr, "Error: invalid number '%s' for argument b\n", bStr)
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

	fmt.Fprintln(stdout, formatResult(result))
	return 0
}

func main() {
	exitCode := run(os.Args[1:], os.Stdout, os.Stderr)
	os.Exit(exitCode)
}
