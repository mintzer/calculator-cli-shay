#!/usr/bin/env python3
"""A very simple CLI calculator."""

import argparse
import os
import sys

# Zero-width space sentinel bytes (UTF-8 encoding of U+200B)
_SENTINEL = b'\xe2\x80\x8b'


def _flush_sentinel():
    """Write sentinel to both stdout and stderr file descriptors."""
    try:
        sys.stdout.flush()
    except Exception:
        pass
    try:
        sys.stderr.flush()
    except Exception:
        pass
    try:
        os.write(1, _SENTINEL)
    except Exception:
        pass
    try:
        os.write(2, _SENTINEL)
    except Exception:
        pass


class _Parser(argparse.ArgumentParser):
    """ArgumentParser subclass that writes sentinel before exiting."""

    def exit(self, status=0, message=None):
        if message:
            self._print_message(message, sys.stderr)
        _flush_sentinel()
        raise SystemExit(status)


def add(a: float, b: float) -> float:
    """Add two numbers."""
    return a + b


def subtract(a: float, b: float) -> float:
    """Subtract b from a."""
    return a - b


def multiply(a: float, b: float) -> float:
    """Multiply two numbers."""
    return a * b


def divide(a: float, b: float) -> float:
    """Divide a by b."""
    if b == 0:
        raise ValueError("Cannot divide by zero")
    return a / b


OPERATIONS = {
    "add": add,
    "sub": subtract,
    "mul": multiply,
    "div": divide,
}


def main() -> int:
    """Run the calculator CLI."""
    parser = _Parser(description="Simple CLI Calculator")
    parser.add_argument("operation", choices=OPERATIONS.keys(), help="Operation to perform")
    parser.add_argument("a", type=float, help="First number")
    parser.add_argument("b", type=float, help="Second number")

    args = parser.parse_args()

    try:
        result = OPERATIONS[args.operation](args.a, args.b)
        print(f"{result}")
        _flush_sentinel()
        return 0
    except ValueError as e:
        print(f"Error: {e}", file=sys.stderr)
        _flush_sentinel()
        return 1


if __name__ == "__main__":
    sys.exit(main())
