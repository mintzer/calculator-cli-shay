# Calculator CLI (Go)

A simple CLI calculator written in Go, migrated from the original Python implementation.

The binary is named `calc-go` to distinguish it from the original Python `calc` tool.

## Build

Requires Go 1.22 or later.

```bash
go build -o calc-go .
```

## Usage

```bash
calc-go <operation> <a> <b>
```

### Operations

| Operation | Description        | Example                  | Output |
|-----------|--------------------|--------------------------|--------|
| `add`     | Addition           | `calc-go add 5 3`       | `8.0`  |
| `sub`     | Subtraction (a-b)  | `calc-go sub 10 4`      | `6.0`  |
| `mul`     | Multiplication     | `calc-go mul 6 7`       | `42.0` |
| `div`     | Division (a/b)     | `calc-go div 20 4`      | `5.0`  |

### Examples

```bash
# Basic arithmetic
./calc-go add 5 3      # 8.0
./calc-go sub 10 4     # 6.0
./calc-go mul 6 7      # 42.0
./calc-go div 20 4     # 5.0

# Float inputs
./calc-go mul 5.5 2    # 11.0
./calc-go div 10 3     # 3.3333333333333335

# Help
./calc-go -h
./calc-go --help
```

### Error Handling

- **Invalid operation**: prints error to stderr, exits with code 2
- **Non-numeric arguments**: prints error to stderr, exits with code 2
- **Missing/extra arguments**: prints usage to stderr, exits with code 2
- **Divide by zero**: prints `Error: Cannot divide by zero` to stderr, exits with code 1

## Development

### Run tests

```bash
go test ./...
```

### Run tests with verbose output

```bash
go test -v ./...
```

## Project Structure

```
├── main.go              # CLI entry point, argument parsing, output formatting
├── calculator.go        # Pure arithmetic functions (Add, Sub, Mul, Div)
├── main_test.go         # CLI integration tests
├── calculator_test.go   # Unit tests for arithmetic functions
├── go.mod               # Go module definition
└── README.md            # This file
```
