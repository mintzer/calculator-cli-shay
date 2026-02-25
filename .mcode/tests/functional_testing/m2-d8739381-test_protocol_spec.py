#!/usr/bin/env python3
"""
BE Testing - CLI Contract Validation Tests

Generated pytest script to validate the CLI spec against the running application.
Each command is tested as a parameterized test case using pytest.

This script supports two modes:
1. SRC Validation: Tests commands and captures outputs (no expected_stdout/stderr)
2. DST Contract Validation: Tests commands and validates outputs match expected

Generated at: 2026-02-25T17:52:40.903288+00:00
Project: calculator-cli-shay
Milestone: 2
"""

import json
import os
import re
import shutil
import subprocess
import sys
import tempfile
import time
from pathlib import Path
from typing import Any

import pytest

# =============================================================================
# Test Configuration (embedded from spec validation)
# =============================================================================

# Parse JSON at runtime to handle null -> None, true -> True, false -> False
TEST_CASES = json.loads(r'''[
    {
        "name": "test_add_positive_integers",
        "category": "HAPPY_PATH",
        "description": "Addition of two positive integers produces correct result",
        "command": ".venv/bin/calc",
        "subcommand": "add",
        "args": [
            "5",
            "3"
        ],
        "expected_exit_code": 0,
        "expected_stdout": "8.0",
        "expected_stderr": null,
        "timeout_seconds": 10
    },
    {
        "name": "test_sub_positive_integers",
        "category": "HAPPY_PATH",
        "description": "Subtraction of two positive integers produces correct result",
        "command": ".venv/bin/calc",
        "subcommand": "sub",
        "args": [
            "10",
            "4"
        ],
        "expected_exit_code": 0,
        "expected_stdout": "6.0",
        "expected_stderr": null,
        "timeout_seconds": 10
    },
    {
        "name": "test_mul_positive_integers",
        "category": "HAPPY_PATH",
        "description": "Multiplication of two positive integers produces correct result",
        "command": ".venv/bin/calc",
        "subcommand": "mul",
        "args": [
            "6",
            "7"
        ],
        "expected_exit_code": 0,
        "expected_stdout": "42.0",
        "expected_stderr": null,
        "timeout_seconds": 10
    },
    {
        "name": "test_div_positive_integers",
        "category": "HAPPY_PATH",
        "description": "Division of two positive integers produces correct result",
        "command": ".venv/bin/calc",
        "subcommand": "div",
        "args": [
            "20",
            "4"
        ],
        "expected_exit_code": 0,
        "expected_stdout": "5.0",
        "expected_stderr": null,
        "timeout_seconds": 10
    },
    {
        "name": "test_add_negative_numbers",
        "category": "HAPPY_PATH",
        "description": "Addition with negative numbers works correctly",
        "command": ".venv/bin/calc",
        "subcommand": "add",
        "args": [
            "-5",
            "3"
        ],
        "expected_exit_code": 0,
        "expected_stdout": "-2.0",
        "expected_stderr": null,
        "timeout_seconds": 10
    },
    {
        "name": "test_sub_negative_numbers",
        "category": "HAPPY_PATH",
        "description": "Subtraction of two negative numbers works correctly",
        "command": ".venv/bin/calc",
        "subcommand": "sub",
        "args": [
            "-10",
            "-4"
        ],
        "expected_exit_code": 0,
        "expected_stdout": "-6.0",
        "expected_stderr": null,
        "timeout_seconds": 10
    },
    {
        "name": "test_mul_negative_numbers",
        "category": "HAPPY_PATH",
        "description": "Multiplication with a negative number produces correct negative result",
        "command": ".venv/bin/calc",
        "subcommand": "mul",
        "args": [
            "-6",
            "7"
        ],
        "expected_exit_code": 0,
        "expected_stdout": "-42.0",
        "expected_stderr": null,
        "timeout_seconds": 10
    },
    {
        "name": "test_div_negative_numbers",
        "category": "HAPPY_PATH",
        "description": "Division with a negative dividend produces correct negative result",
        "command": ".venv/bin/calc",
        "subcommand": "div",
        "args": [
            "-20",
            "4"
        ],
        "expected_exit_code": 0,
        "expected_stdout": "-5.0",
        "expected_stderr": null,
        "timeout_seconds": 10
    },
    {
        "name": "test_add_floating_point",
        "category": "HAPPY_PATH",
        "description": "Addition of floating-point numbers produces correct result",
        "command": ".venv/bin/calc",
        "subcommand": "add",
        "args": [
            "1.5",
            "2.5"
        ],
        "expected_exit_code": 0,
        "expected_stdout": "4.0",
        "expected_stderr": null,
        "timeout_seconds": 10
    },
    {
        "name": "test_div_non_integer_result",
        "category": "HAPPY_PATH",
        "description": "Division producing a floating-point result displays correctly",
        "command": ".venv/bin/calc",
        "subcommand": "div",
        "args": [
            "7",
            "2"
        ],
        "expected_exit_code": 0,
        "expected_stdout": "3.5",
        "expected_stderr": null,
        "timeout_seconds": 10
    },
    {
        "name": "test_add_zeros",
        "category": "BOUNDARY",
        "description": "Addition with zeros produces correct result",
        "command": ".venv/bin/calc",
        "subcommand": "add",
        "args": [
            "0",
            "0"
        ],
        "expected_exit_code": 0,
        "expected_stdout": "0.0",
        "expected_stderr": null,
        "timeout_seconds": 10
    },
    {
        "name": "test_mul_by_zero",
        "category": "BOUNDARY",
        "description": "Multiplication by zero produces zero",
        "command": ".venv/bin/calc",
        "subcommand": "mul",
        "args": [
            "100",
            "0"
        ],
        "expected_exit_code": 0,
        "expected_stdout": "0.0",
        "expected_stderr": null,
        "timeout_seconds": 10
    },
    {
        "name": "test_sub_equal_numbers",
        "category": "BOUNDARY",
        "description": "Subtraction of equal numbers produces zero",
        "command": ".venv/bin/calc",
        "subcommand": "sub",
        "args": [
            "42",
            "42"
        ],
        "expected_exit_code": 0,
        "expected_stdout": "0.0",
        "expected_stderr": null,
        "timeout_seconds": 10
    },
    {
        "name": "test_div_one",
        "category": "BOUNDARY",
        "description": "Division by one returns the dividend unchanged",
        "command": ".venv/bin/calc",
        "subcommand": "div",
        "args": [
            "99",
            "1"
        ],
        "expected_exit_code": 0,
        "expected_stdout": "99.0",
        "expected_stderr": null,
        "timeout_seconds": 10
    },
    {
        "name": "test_large_numbers",
        "category": "BOUNDARY",
        "description": "Operations with large numbers work correctly",
        "command": ".venv/bin/calc",
        "subcommand": "add",
        "args": [
            "1000000",
            "2000000"
        ],
        "expected_exit_code": 0,
        "expected_stdout": "3000000.0",
        "expected_stderr": null,
        "timeout_seconds": 10
    },
    {
        "name": "test_very_small_float",
        "category": "BOUNDARY",
        "description": "Operations with very small floating-point numbers work correctly",
        "command": ".venv/bin/calc",
        "subcommand": "mul",
        "args": [
            "0.001",
            "0.001"
        ],
        "expected_exit_code": 0,
        "expected_stdout": "1e-06",
        "expected_stderr": null,
        "timeout_seconds": 10
    },
    {
        "name": "test_div_by_zero",
        "category": "INVALID_ARGS",
        "description": "Division by zero produces an error message and non-zero exit code",
        "command": ".venv/bin/calc",
        "subcommand": "div",
        "args": [
            "1",
            "0"
        ],
        "expected_exit_code": 1,
        "expected_stdout": null,
        "expected_stderr": "Cannot divide by zero",
        "timeout_seconds": 10
    },
    {
        "name": "test_div_zero_by_zero",
        "category": "INVALID_ARGS",
        "description": "Division of zero by zero also produces division-by-zero error",
        "command": ".venv/bin/calc",
        "subcommand": "div",
        "args": [
            "0",
            "0"
        ],
        "expected_exit_code": 1,
        "expected_stdout": null,
        "expected_stderr": "Cannot divide by zero",
        "timeout_seconds": 10
    },
    {
        "name": "test_unknown_operation",
        "category": "INVALID_ARGS",
        "description": "Unknown operation name produces error with invalid choice message",
        "command": ".venv/bin/calc",
        "subcommand": "",
        "args": [
            "foo",
            "1",
            "2"
        ],
        "expected_exit_code": 2,
        "expected_stdout": null,
        "expected_stderr": "invalid choice",
        "timeout_seconds": 10
    },
    {
        "name": "test_unknown_operation_mod",
        "category": "INVALID_ARGS",
        "description": "Unsupported operation 'mod' is rejected",
        "command": ".venv/bin/calc",
        "subcommand": "",
        "args": [
            "mod",
            "10",
            "3"
        ],
        "expected_exit_code": 2,
        "expected_stdout": null,
        "expected_stderr": "invalid choice",
        "timeout_seconds": 10
    },
    {
        "name": "test_missing_all_args",
        "category": "INVALID_ARGS",
        "description": "No arguments at all produces usage text and non-zero exit",
        "command": ".venv/bin/calc",
        "subcommand": "",
        "args": [],
        "expected_exit_code": 2,
        "expected_stdout": null,
        "expected_stderr": "the following arguments are required",
        "timeout_seconds": 10
    },
    {
        "name": "test_missing_operands",
        "category": "INVALID_ARGS",
        "description": "Operation with missing operands produces error and non-zero exit",
        "command": ".venv/bin/calc",
        "subcommand": "add",
        "args": [],
        "expected_exit_code": 2,
        "expected_stdout": null,
        "expected_stderr": "the following arguments are required",
        "timeout_seconds": 10
    },
    {
        "name": "test_missing_second_operand",
        "category": "INVALID_ARGS",
        "description": "Operation with only one operand produces error and non-zero exit",
        "command": ".venv/bin/calc",
        "subcommand": "add",
        "args": [
            "1"
        ],
        "expected_exit_code": 2,
        "expected_stdout": null,
        "expected_stderr": "the following arguments are required",
        "timeout_seconds": 10
    },
    {
        "name": "test_too_many_args",
        "category": "INVALID_ARGS",
        "description": "Extra arguments beyond the expected three produce error and non-zero exit",
        "command": ".venv/bin/calc",
        "subcommand": "add",
        "args": [
            "1",
            "2",
            "3"
        ],
        "expected_exit_code": 2,
        "expected_stdout": null,
        "expected_stderr": "unrecognized arguments",
        "timeout_seconds": 10
    },
    {
        "name": "test_non_numeric_first_operand",
        "category": "INVALID_ARGS",
        "description": "Non-numeric first operand produces an error",
        "command": ".venv/bin/calc",
        "subcommand": "add",
        "args": [
            "abc",
            "2"
        ],
        "expected_exit_code": 2,
        "expected_stdout": null,
        "expected_stderr": "invalid float value",
        "timeout_seconds": 10
    },
    {
        "name": "test_non_numeric_second_operand",
        "category": "INVALID_ARGS",
        "description": "Non-numeric second operand produces an error",
        "command": ".venv/bin/calc",
        "subcommand": "add",
        "args": [
            "1",
            "xyz"
        ],
        "expected_exit_code": 2,
        "expected_stdout": null,
        "expected_stderr": "invalid float value",
        "timeout_seconds": 10
    },
    {
        "name": "test_both_operands_non_numeric",
        "category": "INVALID_ARGS",
        "description": "Both operands non-numeric produces an error for the first",
        "command": ".venv/bin/calc",
        "subcommand": "mul",
        "args": [
            "foo",
            "bar"
        ],
        "expected_exit_code": 2,
        "expected_stdout": null,
        "expected_stderr": "invalid float value",
        "timeout_seconds": 10
    },
    {
        "name": "test_empty_string_operand",
        "category": "BOUNDARY",
        "description": "Empty string as operand is treated as invalid",
        "command": ".venv/bin/calc",
        "subcommand": "add",
        "args": [
            "",
            "5"
        ],
        "expected_exit_code": 2,
        "expected_stdout": null,
        "expected_stderr": "invalid float value",
        "timeout_seconds": 10
    },
    {
        "name": "test_unknown_operation_shows_usage",
        "category": "INVALID_ARGS",
        "description": "Unknown operation error message also includes available choices",
        "command": ".venv/bin/calc",
        "subcommand": "",
        "args": [
            "pow",
            "2",
            "8"
        ],
        "expected_exit_code": 2,
        "expected_stdout": null,
        "expected_stderr": "choose from",
        "timeout_seconds": 10
    }
]''')

# CLI binary/entry point
CLI_COMMAND = ".venv/bin/calc"

# Working directory for CLI execution
WORKING_DIR = "."

# Default command timeout in seconds
DEFAULT_TIMEOUT = 30

# Response validation mode: when True, validates output against expected
VALIDATE_OUTPUT = any(
    tc.get("actual_stdout") is not None or tc.get("actual_stderr") is not None
    for tc in TEST_CASES
)

# =============================================================================
# Output Validation Utilities
# =============================================================================



def normalize_output(output: str) -> str:
    """Normalize output for comparison (strip whitespace, normalize newlines)."""
    if output is None:
        return ""
    return output.strip().replace("\r\n", "\n")


def matches_pattern(actual: str, pattern: str | None) -> bool:
    """
    Check if actual output matches the expected pattern.

    Pattern matching rules:
    - If pattern is None, always matches (no validation)
    - If pattern starts with 'regex:', use regex matching
    - Otherwise, check if pattern is contained in actual output (case-insensitive)
    """
    if pattern is None:
        return True

    actual_normalized = normalize_output(actual)

    if pattern.startswith("regex:"):
        regex_pattern = pattern[6:]  # Remove 'regex:' prefix
        return bool(re.search(regex_pattern, actual_normalized, re.IGNORECASE | re.MULTILINE))

    # Default: substring match (case-insensitive)
    pattern_normalized = normalize_output(pattern)
    return pattern_normalized.lower() in actual_normalized.lower()


def validate_cli_output(
    actual_stdout: str,
    actual_stderr: str,
    expected_stdout: str | None,
    expected_stderr: str | None,
) -> tuple[bool, list[str]]:
    """
    Validate CLI output against expected patterns.

    Args:
        actual_stdout: Actual stdout from command
        actual_stderr: Actual stderr from command
        expected_stdout: Expected stdout pattern (or None)
        expected_stderr: Expected stderr pattern (or None)

    Returns:
        tuple: (is_valid, list of violations)
    """
    violations: list[str] = []

    if expected_stdout is not None and not matches_pattern(actual_stdout, expected_stdout):
        violations.append(
            f"stdout mismatch: expected pattern '{expected_stdout}' not found in output"
        )

    if expected_stderr is not None and not matches_pattern(actual_stderr, expected_stderr):
        violations.append(
            f"stderr mismatch: expected pattern '{expected_stderr}' not found in output"
        )

    return len(violations) == 0, violations


def format_output_diff(violations: list[str]) -> str:
    """Format output differences for error message."""
    if not violations:
        return "No differences"

    output = []
    for i, diff in enumerate(violations):
        output.append(f"  - {diff}")

    return "\n".join(output)


# =============================================================================
# Test Results Collection
# =============================================================================

test_results: list[dict[str, Any]] = []


def record_result(
    name: str,
    command: str,
    subcommand: str | None,
    args: list[str],
    expected_exit_code: int,
    actual_exit_code: int,
    passed: bool,
    duration_ms: float,
    category: str | None = None,
    description: str | None = None,
    error: str | None = None,
    stdout: str | None = None,
    stderr: str | None = None,
    output_match: bool | None = None,
    output_diff: list[str] | None = None,
) -> None:
    """Record a test result for final output."""
    result: dict[str, Any] = {
        "name": name,
        "command": command,
        "subcommand": subcommand,
        "args": args,
        "expected_exit_code": expected_exit_code,
        "actual_exit_code": actual_exit_code,
        "passed": passed,
        "duration_ms": duration_ms,
        "category": category,
        "description": description,
    }
    if error:
        result["error"] = error

    # Track output validation results (for DST contract testing)
    if output_match is not None:
        result["output_match"] = output_match
    if output_diff:
        result["output_diff"] = output_diff

    # Capture outputs for validation
    if stdout:
        if passed:
            result["actual_stdout"] = stdout  # Capture more for passed tests
        else:
            result["stdout"] = stdout

    if stderr:
        if passed:
            result["actual_stderr"] = stderr
        else:
            result["stderr"] = stderr

    test_results.append(result)


# =============================================================================
# Setup and Cleanup Helpers
# =============================================================================


def run_setup(setup_config: dict[str, Any], work_dir: Path) -> bool:
    """Run setup actions before a test."""
    if not setup_config:
        return True

    try:
        # Create file
        if "create_file" in setup_config:
            file_config = setup_config["create_file"]
            file_path = work_dir / file_config["path"]
            file_path.parent.mkdir(parents=True, exist_ok=True)
            file_path.write_text(file_config.get("content", ""))
            print(f"Setup: Created file {file_path}")

        # Create directory
        if "create_dir" in setup_config:
            dir_path = work_dir / setup_config["create_dir"]
            dir_path.mkdir(parents=True, exist_ok=True)
            print(f"Setup: Created directory {dir_path}")

        # Run command
        if "run_command" in setup_config:
            cmd = setup_config["run_command"]
            result = subprocess.run(
                cmd,
                shell=True,
                cwd=str(work_dir),
                capture_output=True,
                text=True,
                timeout=DEFAULT_TIMEOUT,
            )
            if result.returncode != 0:
                print(f"Setup command failed: {result.stderr}")
                return False

        return True

    except Exception as e:
        print(f"Setup error: {e}")
        return False


def run_cleanup(cleanup_config: dict[str, Any], work_dir: Path) -> None:
    """Run cleanup actions after a test (best effort)."""
    if not cleanup_config:
        return

    try:
        # Delete files
        if "delete_files" in cleanup_config:
            for file_path in cleanup_config["delete_files"]:
                full_path = work_dir / file_path
                if full_path.exists():
                    full_path.unlink()
                    print(f"Cleanup: Deleted file {full_path}")

        # Delete directories
        if "delete_dirs" in cleanup_config:
            for dir_path in cleanup_config["delete_dirs"]:
                full_path = work_dir / dir_path
                if full_path.exists():
                    shutil.rmtree(full_path)
                    print(f"Cleanup: Deleted directory {full_path}")

        # Run command
        if "run_command" in cleanup_config:
            cmd = cleanup_config["run_command"]
            subprocess.run(
                cmd,
                shell=True,
                cwd=str(work_dir),
                capture_output=True,
                text=True,
                timeout=DEFAULT_TIMEOUT,
            )

    except Exception as e:
        print(f"Cleanup warning: {e}")


# =============================================================================
# Pytest Fixtures
# =============================================================================


@pytest.fixture(scope="session")
def cli_work_dir() -> Path:
    """Get the CLI working directory."""
    return Path(WORKING_DIR)


@pytest.fixture(scope="session", autouse=True)
def verify_cli_exists() -> None:
    """Verify the CLI command exists before running tests."""
    print(f"\nVerifying CLI command exists: {CLI_COMMAND}...")

    # Check if it's a direct path
    if os.path.isfile(CLI_COMMAND):
        print(f"CLI found at: {CLI_COMMAND}")
        return

    # Check if it's in PATH
    result = shutil.which(CLI_COMMAND)
    if result:
        print(f"CLI found in PATH: {result}")
        return

    # Try common locations
    work_dir = Path(WORKING_DIR)
    common_paths = [
        work_dir / CLI_COMMAND,
        work_dir / "dist" / CLI_COMMAND,
        work_dir / "target" / "release" / CLI_COMMAND,
        work_dir / "bin" / CLI_COMMAND,
    ]

    for path in common_paths:
        if path.exists():
            print(f"CLI found at: {path}")
            return

    pytest.fail(f"CLI command '{CLI_COMMAND}' not found. Please ensure the app is built.")


# =============================================================================
# Test Cases
# =============================================================================


def get_test_ids() -> list[str]:
    """Generate test IDs for parametrization."""
    return [tc.get("name", f"test_{i}") for i, tc in enumerate(TEST_CASES)]


@pytest.mark.parametrize("test_case", TEST_CASES, ids=get_test_ids())
def test_cli_command(test_case: dict[str, Any], cli_work_dir: Path) -> None:
    """Test a single CLI command based on test case configuration."""
    # Extract test case info
    name = test_case.get("name", "unnamed")
    command = CLI_COMMAND
    raw_args = test_case.get("args", [])
    args = (
        [str(arg) for arg in raw_args]
        if isinstance(raw_args, list)
        else ([str(raw_args)] if raw_args is not None else [])
    )
    subcommand = test_case.get("subcommand", "")
    subcommand_parts = (
        [part for part in subcommand.strip().split(" ") if part]
        if isinstance(subcommand, str) and subcommand.strip()
        else []
    )
    execution_args = subcommand_parts + args
    stdin_input = test_case.get("stdin")
    env_vars = test_case.get("env", {})
    expected_exit_code = test_case.get("expected_exit_code", 0)
    expected_stdout = test_case.get("expected_stdout")
    expected_stderr = test_case.get("expected_stderr")
    category = test_case.get("category")
    description = test_case.get("description")
    setup_config = test_case.get("setup")
    cleanup_config = test_case.get("cleanup")
    timeout = test_case.get("timeout_seconds", DEFAULT_TIMEOUT)

    # Expected outputs for DST contract validation (from SRC validation)
    actual_stdout_expected = test_case.get("actual_stdout")
    actual_stderr_expected = test_case.get("actual_stderr")

    try:
        # Run setup if configured
        if setup_config:
            if not run_setup(setup_config, cli_work_dir):
                record_result(
                    name=name,
                    command=command,
                    subcommand=subcommand if isinstance(subcommand, str) and subcommand.strip() else None,
                    args=execution_args,
                    expected_exit_code=expected_exit_code,
                    actual_exit_code=-1,
                    passed=False,
                    duration_ms=0,
                    category=category,
                    description=description,
                    error="Setup failed",
                )
                pytest.fail(f"Setup failed for test '{name}'")

        # Build full command
        full_cmd = [command] + execution_args

        # Prepare environment
        env = os.environ.copy()
        env.update(env_vars)

        # Execute command
        start_time = time.time()
        try:
            result = subprocess.run(
                full_cmd,
                input=stdin_input,
                cwd=str(cli_work_dir),
                env=env,
                capture_output=True,
                text=True,
                timeout=timeout,
            )

            duration_ms = (time.time() - start_time) * 1000
            actual_exit_code = result.returncode
            stdout = result.stdout
            stderr = result.stderr

            # Check exit code first
            exit_code_passed = actual_exit_code == expected_exit_code
            error_msg = None if exit_code_passed else (
                f"Expected exit code {expected_exit_code}, got {actual_exit_code}"
            )

            # Check output patterns
            output_match: bool | None = None
            output_diff: list[str] | None = None

            # For DST validation, compare against captured SRC output
            if actual_stdout_expected is not None or actual_stderr_expected is not None:
                output_match, output_diff = validate_cli_output(
                    stdout,
                    stderr,
                    actual_stdout_expected,
                    actual_stderr_expected,
                )
                if not output_match:
                    error_msg = f"Output contract violation:\n{format_output_diff(output_diff)}"
            # For SRC validation or basic validation, check expected patterns
            elif expected_stdout is not None or expected_stderr is not None:
                output_match, output_diff = validate_cli_output(
                    stdout,
                    stderr,
                    expected_stdout,
                    expected_stderr,
                )
                if not output_match:
                    error_msg = f"Output pattern mismatch:\n{format_output_diff(output_diff)}"

            # Overall pass
            passed = exit_code_passed and (output_match is None or output_match)

            record_result(
                name=name,
                command=command,
                subcommand=subcommand if isinstance(subcommand, str) and subcommand.strip() else None,
                args=execution_args,
                expected_exit_code=expected_exit_code,
                actual_exit_code=actual_exit_code,
                passed=passed,
                duration_ms=duration_ms,
                category=category,
                description=description,
                error=error_msg,
                stdout=stdout,
                stderr=stderr,
                output_match=output_match,
                output_diff=output_diff,
            )

            # pytest assertions
            if not exit_code_passed:
                pytest.fail(
                    f"Test '{name}': Expected exit code {expected_exit_code}, got {actual_exit_code}.\n"
                    f"stdout: {stdout if stdout else 'empty'}\n"
                    f"stderr: {stderr if stderr else 'empty'}"
                )

            if output_match is False:
                pytest.fail(
                    f"Test '{name}': Output validation failed.\n"
                    f"Violations:\n{format_output_diff(output_diff or [])}"
                )

        except subprocess.TimeoutExpired as e:
            duration_ms = (time.time() - start_time) * 1000
            record_result(
                name=name,
                command=command,
                subcommand=subcommand if isinstance(subcommand, str) and subcommand.strip() else None,
                args=execution_args,
                expected_exit_code=expected_exit_code,
                actual_exit_code=-1,
                passed=False,
                duration_ms=duration_ms,
                category=category,
                description=description,
                error=f"Command timed out after {timeout}s",
                stdout=e.stdout if hasattr(e, 'stdout') else None,
                stderr=e.stderr if hasattr(e, 'stderr') else None,
            )
            pytest.fail(f"Test '{name}': Command timed out after {timeout}s")

    except Exception as e:
        record_result(
            name=name,
            command=command,
            subcommand=subcommand if isinstance(subcommand, str) and subcommand.strip() else None,
            args=execution_args,
            expected_exit_code=expected_exit_code,
            actual_exit_code=-1,
            passed=False,
            duration_ms=0,
            category=category,
            description=description,
            error=f"Test error: {type(e).__name__}: {e}",
        )
        raise

    finally:
        # Always run cleanup
        if cleanup_config:
            run_cleanup(cleanup_config, cli_work_dir)


# =============================================================================
# Test Results Output
# =============================================================================


@pytest.fixture(scope="session", autouse=True)
def output_test_results(request: pytest.FixtureRequest) -> Any:
    """Output test results in JSON format after all tests complete."""
    yield  # Wait for all tests to complete

    # Calculate final results
    passed_count = sum(1 for r in test_results if r["passed"])
    failed_count = len([r for r in test_results if not r["passed"]])
    total_count = len(test_results)
    all_passed = failed_count == 0 and total_count > 0

    failures = [r for r in test_results if not r["passed"]]

    # Count output validation results (for DST contract testing)
    output_validated_count = sum(1 for r in test_results if r.get("output_match") is not None)
    output_match_count = sum(1 for r in test_results if r.get("output_match") is True)

    output = {
        "all_passed": all_passed,
        "passed_count": passed_count,
        "failed_count": failed_count,
        "total_count": total_count,
        "results": test_results,
        "failures": failures,
    }

    # Add contract validation summary if any tests had expected outputs
    if output_validated_count > 0:
        output["contract_validation"] = {
            "tests_with_expected_output": output_validated_count,
            "output_matches": output_match_count,
            "output_mismatches": output_validated_count - output_match_count,
        }

    print("\n" + "=" * 60)
    print(f"Results: {passed_count}/{total_count} passed")
    if output_validated_count > 0:
        print(f"Contract validation: {output_match_count}/{output_validated_count} outputs matched")
    print("=" * 60)
    print(json.dumps(output))
    sys.stdout.flush()
