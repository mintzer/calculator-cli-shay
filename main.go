package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

type opFunc func(a, b float64) (float64, error)

var operations = map[string]opFunc{
	"add": func(a, b float64) (float64, error) { return add(a, b), nil },
	"sub": func(a, b float64) (float64, error) { return subtract(a, b), nil },
	"mul": func(a, b float64) (float64, error) { return multiply(a, b), nil },
	"div": func(a, b float64) (float64, error) { return divide(a, b) },
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: calc <operation> <a> <b>")
	fmt.Fprintln(os.Stderr, "Operations: add, sub, mul, div")
}

func main() {
	if len(os.Args) != 4 {
		usage()
		os.Exit(1)
	}

	opName := os.Args[1]
	aStr := os.Args[2]
	bStr := os.Args[3]

	op, ok := operations[opName]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: Unknown operation %q\n", opName)
		usage()
		os.Exit(1)
	}

	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Invalid number %q\n", aStr)
		os.Exit(1)
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Invalid number %q\n", bStr)
		os.Exit(1)
	}

	result, err := op(a, b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
