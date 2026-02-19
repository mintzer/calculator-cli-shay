#!/usr/bin/env python3
"""A very simple CLI calculator (Go migration - Python bridge)."""

import argparse
import sys


class CompatArgumentParser(argparse.ArgumentParser):
    """ArgumentParser that ensures consistent choice formatting across Python versions."""

    def _check_value(self, action, value):
        """Override to ensure choices are printed without quotes (Python 3.12.8+ style)."""
        if action.choices is not None and value not in action.choices:
            args = {
                "value": value,
                "choices": ", ".join([str(c) for c in action.choices]),
            }
            msg = "invalid choice: %(value)r (choose from %(choices)s)" % args
            raise argparse.ArgumentError(action, msg)


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


def main():
    """Run the calculator CLI."""
    parser = CompatArgumentParser(prog="calc", description="Simple CLI Calculator")
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
