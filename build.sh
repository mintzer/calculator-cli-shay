#!/bin/sh
# Build script that selects the pre-compiled binary for the current platform.
# Falls back to `go build` if Go is available.

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Map architecture names
case "$ARCH" in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
esac

BINARY="calc-go-${OS}-${ARCH}"

if [ -f "$BINARY" ]; then
    cp "$BINARY" calc-go
    chmod +x calc-go
    echo "Using pre-compiled binary: $BINARY"
elif command -v go >/dev/null 2>&1; then
    go build -o calc-go .
    echo "Built with go build"
else
    echo "Error: No pre-compiled binary for ${OS}/${ARCH} and Go is not installed" >&2
    exit 1
fi
