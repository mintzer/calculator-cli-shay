"""CLI integration tests for the calculator."""

import subprocess
import sys
from pathlib import Path

import pytest

from calculator import main

PROJECT_ROOT = Path(__file__).resolve().parent.parent


class TestCLIHappyPath:
    """Test successful CLI operations using main() directly."""

    def test_add(self, monkeypatch, capsys):
        monkeypatch.setattr(sys, "argv", ["calc_3_12", "add", "5", "3"])
        exit_code = main()
        captured = capsys.readouterr()
        assert exit_code == 0
        assert captured.out.strip() == "8.0"
        assert captured.err == ""

    def test_sub(self, monkeypatch, capsys):
        monkeypatch.setattr(sys, "argv", ["calc_3_12", "sub", "5", "3"])
        exit_code = main()
        captured = capsys.readouterr()
        assert exit_code == 0
        assert captured.out.strip() == "2.0"
        assert captured.err == ""

    def test_mul(self, monkeypatch, capsys):
        monkeypatch.setattr(sys, "argv", ["calc_3_12", "mul", "2", "4"])
        exit_code = main()
        captured = capsys.readouterr()
        assert exit_code == 0
        assert captured.out.strip() == "8.0"
        assert captured.err == ""

    def test_div(self, monkeypatch, capsys):
        monkeypatch.setattr(sys, "argv", ["calc_3_12", "div", "8", "4"])
        exit_code = main()
        captured = capsys.readouterr()
        assert exit_code == 0
        assert captured.out.strip() == "2.0"
        assert captured.err == ""

    def test_negative_numbers(self, monkeypatch, capsys):
        monkeypatch.setattr(sys, "argv", ["calc_3_12", "add", "-5", "3"])
        exit_code = main()
        captured = capsys.readouterr()
        assert exit_code == 0
        assert captured.out.strip() == "-2.0"

    def test_float_inputs(self, monkeypatch, capsys):
        monkeypatch.setattr(sys, "argv", ["calc_3_12", "mul", "2.5", "4"])
        exit_code = main()
        captured = capsys.readouterr()
        assert exit_code == 0
        assert captured.out.strip() == "10.0"


class TestCLIErrorPaths:
    """Test error handling in CLI."""

    def test_divide_by_zero(self, monkeypatch, capsys):
        monkeypatch.setattr(sys, "argv", ["calc_3_12", "div", "5", "0"])
        exit_code = main()
        captured = capsys.readouterr()
        assert exit_code == 1
        assert "Cannot divide by zero" in captured.err
        assert captured.out == ""

    def test_invalid_operation(self, monkeypatch):
        monkeypatch.setattr(sys, "argv", ["calc_3_12", "pow", "2", "3"])
        with pytest.raises(SystemExit) as exc_info:
            main()
        assert exc_info.value.code != 0


class TestCLISubprocess:
    """End-to-end tests using subprocess to verify the actual script."""

    def test_add_subprocess(self):
        result = subprocess.run(
            [sys.executable, "calculator.py", "add", "5", "3"],
            capture_output=True,
            text=True,
            cwd=PROJECT_ROOT,
        )
        assert result.returncode == 0
        assert result.stdout.strip() == "8.0"

    def test_div_by_zero_subprocess(self):
        result = subprocess.run(
            [sys.executable, "calculator.py", "div", "5", "0"],
            capture_output=True,
            text=True,
            cwd=PROJECT_ROOT,
        )
        assert result.returncode == 1
        assert "Cannot divide by zero" in result.stderr

    def test_invalid_operation_subprocess(self):
        result = subprocess.run(
            [sys.executable, "calculator.py", "pow", "2", "3"],
            capture_output=True,
            text=True,
            cwd=PROJECT_ROOT,
        )
        assert result.returncode != 0
        assert "invalid choice" in result.stderr
