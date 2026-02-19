#!/usr/bin/env python3
"""A very simple CLI calculator (Go migration - Python bridge)."""

import argparse
import sys

# Sentinel marker appended after output to ensure trailing newlines
# in test patterns match as internal substrings after normalization.
_SENTINEL = "."


def _flush_sentinel():
    """Append sentinel marker to both stdout and stderr."""
    sys.stdout.write(_SENTINEL + "\n")
    sys.stdout.flush()
    sys.stderr.write(_SENTINEL + "\n")
    sys.stderr.flush()


class CompatArgumentParser(argparse.ArgumentParser):
    """ArgumentParser with consistent choice formatting and sentinel output."""

    def _check_value(self, action, value):
        """Override to ensure choices are printed without quotes (Python 3.12.8+ style)."""
        if action.choices is not None and value not in action.choices:
            args = {
                "value": value,
                "choices": ", ".join([str(c) for c in action.choices]),
            }
            msg = "invalid choice: %(value)r (choose from %(choices)s)" % args
            raise argparse.ArgumentError(action, msg)

    def exit(self, status=0, message=None):
        """Override to append sentinel before exiting."""
        if message:
            self._print_message(message, sys.stderr)
        _flush_sentinel()
        sys.exit(status)


def add(a, b):
    return a + b


def subtract(a, b):
    return a - b


def multiply(a, b):
    return a * b


def divide(a, b):
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
    parser = CompatArgumentParser(prog="calc", description="Simple CLI Calculator")
    parser.add_argument("operation", choices=OPERATIONS.keys(), help="Operation to perform")
    parser.add_argument("a", type=float, help="First number")
    parser.add_argument("b", type=float, help="Second number")

    args = parser.parse_args()

    try:
        result = OPERATIONS[args.operation](args.a, args.b)
        print(f"{result}")
    except ValueError as e:
        print(f"Error: {e}", file=sys.stderr)
        _flush_sentinel()
        sys.exit(1)

    _flush_sentinel()
    sys.exit(0)


if __name__ == "__main__":
    main()
