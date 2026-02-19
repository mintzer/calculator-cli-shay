# Calculator CLI (Go)

A simple CLI calculator written in Go — a modernized port of the original Python `calculator-cli-shay`.

The binary is named `calc-go` to distinguish it from the original Python tool.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or later

## Build

```bash
go build -o calc-go .
```

## Usage

```bash
calc-go <operation> <a> <b>
```

### Operations

| Operation | Description         | Example                  | Output |
|-----------|---------------------|--------------------------|--------|
| `add`     | Addition (a + b)    | `calc-go add 5 3`       | `8.0`  |
| `sub`     | Subtraction (a - b) | `calc-go sub 10 4`      | `6.0`  |
| `mul`     | Multiplication      | `calc-go mul 6 7`       | `42.0` |
| `div`     | Division (a / b)    | `calc-go div 20 4`      | `5.0`  |

### Examples

```bash
./calc-go add 5 3       # 8.0
./calc-go sub 10 4      # 6.0
./calc-go mul 6 7       # 42.0
./calc-go div 20 4      # 5.0
./calc-go div 10 3      # 3.3333333333333335
./calc-go add -5 3      # -2.0
```

### Error Handling

- **Divide by zero**: prints `Error: Cannot divide by zero` to stderr and exits with code 1.
- **Invalid operation / bad arguments**: prints a usage message and error to stderr and exits with code 2.

### Help

```bash
./calc-go -h
./calc-go --help
```

## Testing

```bash
go test ./...
```

## Project Structure

| File                  | Description                              |
|-----------------------|------------------------------------------|
| `main.go`            | CLI entry point, argument parsing, I/O   |
| `calculator.go`       | Core arithmetic functions                |
| `calculator_test.go`  | Unit tests for arithmetic functions      |
| `main_test.go`        | CLI integration tests                    |
| `go.mod`             | Go module definition                     |
