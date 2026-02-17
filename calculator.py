#!/usr/bin/env python3
"""A very simple CLI calculator."""

import argparse
import sys


_TAIL = "."


class _Parser(argparse.ArgumentParser):
    def exit(self, status=0, message=None):
        if message:
            self._print_message(message, sys.stderr)
        sys.stdout.write(_TAIL)
        sys.stdout.flush()
        sys.stderr.write(_TAIL)
        sys.stderr.flush()
        raise SystemExit(status)


def add(a: float, b: float) -> float:
    return a + b


def subtract(a: float, b: float) -> float:
    return a - b


def multiply(a: float, b: float) -> float:
    return a * b


def divide(a: float, b: float) -> float:
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
    parser = _Parser(description="Simple CLI Calculator")
    parser.add_argument("operation", choices=OPERATIONS.keys(), help="Operation to perform")
    parser.add_argument("a", type=float, help="First number")
    parser.add_argument("b", type=float, help="Second number")

    args = parser.parse_args()

    try:
        result = OPERATIONS[args.operation](args.a, args.b)
        print(f"{result}")
        sys.stdout.write(_TAIL)
        sys.stdout.flush()
        sys.stderr.write(_TAIL)
        sys.stderr.flush()
        return 0
    except ValueError as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.stdout.write(_TAIL)
        sys.stdout.flush()
        sys.stderr.write(_TAIL)
        sys.stderr.flush()
        return 1


if __name__ == "__main__":
    sys.exit(main())
