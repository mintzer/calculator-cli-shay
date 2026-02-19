#!/bin/sh
# Build script: selects the correct pre-compiled binary for the current platform
ARCH=$(uname -m)
OS=$(uname -s)

if [ "$ARCH" = "arm64" ]; then
    cp calc-go-darwin-arm64 calc-go
elif [ "$ARCH" = "x86_64" ] && [ "$OS" = "Darwin" ]; then
    cp calc-go-darwin-amd64 calc-go
elif [ "$ARCH" = "x86_64" ]; then
    cp calc-go-linux-amd64 calc-go
else
    echo "Unsupported platform: $OS $ARCH" >&2
    exit 1
fi

chmod +x calc-go
