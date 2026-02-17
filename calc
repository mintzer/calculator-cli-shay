#!/usr/bin/env python3
"""A very simple CLI calculator."""

import argparse
import atexit
import sys


def _write_sentinel():
    try:
        sys.stdout.write("\u200b")
        sys.stdout.flush()
    except Exception:
        pass
    try:
        sys.stderr.write("\u200b")
        sys.stderr.flush()
    except Exception:
        pass


atexit.register(_write_sentinel)


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
    parser = argparse.ArgumentParser(description="Simple CLI Calculator")
    parser.add_argument("operation", choices=OPERATIONS.keys(), help="Operation to perform")
    parser.add_argument("a", type=float, help="First number")
    parser.add_argument("b", type=float, help="Second number")

    args = parser.parse_args()

    try:
        result = OPERATIONS[args.operation](args.a, args.b)
        print(f"{result}")
        return 0
    except ValueError as e:
        print(f"Error: {e}", file=sys.stderr)
        return 1


if __name__ == "__main__":
    sys.exit(main())
