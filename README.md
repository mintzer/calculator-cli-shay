# Calculator CLI (Go)

A simple CLI calculator written in Go — a modernized port of the original Python `calculator-cli-shay`.

The binary is named `calc-go` to distinguish it from the original Python version.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or later

## Build

```bash
go build -o calc-go .
```

## Usage

```bash
./calc-go <operation> <a> <b>
```

### Operations

| Operation | Description        | Example                  | Output |
|-----------|--------------------|--------------------------|--------|
| `add`     | Add two numbers    | `./calc-go add 5 3`     | `8.0`  |
| `sub`     | Subtract b from a  | `./calc-go sub 10 4`    | `6.0`  |
| `mul`     | Multiply two numbers | `./calc-go mul 6 7`   | `42.0` |
| `div`     | Divide a by b      | `./calc-go div 20 4`    | `5.0`  |

### Help

```bash
./calc-go -h
./calc-go --help
```

### Error Handling

- **Invalid operation**: Prints usage and error to stderr, exits with code 2.
- **Non-numeric arguments**: Prints usage and error to stderr, exits with code 2.
- **Missing arguments**: Prints usage and error to stderr, exits with code 2.
- **Divide by zero**: Prints `Error: Cannot divide by zero` to stderr, exits with code 1.

## Testing

```bash
go test ./...
```

Run with verbose output:

```bash
go test -v ./...
```
