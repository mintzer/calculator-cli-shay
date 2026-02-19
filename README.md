# calc-go

A simple CLI calculator — Go port of [calculator-cli-shay](https://github.com/mintzer/calculator-cli-shay).

## Build

```sh
go build -o calc-go .
```

## Usage

```sh
calc-go <operation> <a> <b>
```

### Operations

| Operation | Description          |
|-----------|----------------------|
| `add`     | Add two numbers      |
| `sub`     | Subtract b from a    |
| `mul`     | Multiply two numbers |
| `div`     | Divide a by b        |

### Examples

```sh
./calc-go add 5 3      # 8.0
./calc-go sub 10 4     # 6.0
./calc-go mul 3 7      # 21.0
./calc-go div 10 3     # 3.3333333333333335
```

### Help

```sh
./calc-go --help
```

## Testing

```sh
go test ./...
```
