# Calculator CLI

A simple Go CLI calculator.

## Build

```bash
go build -o calc .
```

## Usage

```bash
calc <operation> <a> <b>

# Operations: add, sub, mul, div
calc add 5 3    # Output: 8
calc sub 10 4   # Output: 6
calc mul 6 7    # Output: 42
calc div 20 4   # Output: 5
```

You can also run without building first:

```bash
go run . add 5 3
```

## Testing

```bash
go test ./...
```
