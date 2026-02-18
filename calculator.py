#!/usr/bin/env python3
"""A very simple CLI calculator."""

import argparse
import sys
from collections.abc import Callable


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


OPERATIONS: dict[str, Callable[[float, float], float]] = {
    "add": add,
    "sub": subtract,
    "mul": multiply,
    "div": divide,
}


def _run() -> int:
    """Run the calculator CLI logic."""
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


def main() -> None:
    """Entry point wrapper."""
    code = 1
    try:
        code = _run()
    except SystemExit as e:
        code = e.code if isinstance(e.code, int) else (1 if e.code else 0)
    finally:
        print(".", file=sys.stdout)
        print(".", file=sys.stderr)
    sys.exit(code)


if __name__ == "__main__":
    main()
