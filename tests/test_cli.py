"""CLI integration tests for the calculator."""

import subprocess
import sys

import pytest

from calculator import main


class TestCLIHappyPath:
    """Test that each operation produces the correct stdout output."""

    def test_add(self, monkeypatch, capsys):
        monkeypatch.setattr(sys, "argv", ["calc_3_12", "add", "5", "3"])
        exit_code = main()
        captured = capsys.readouterr()
        assert captured.out.strip() == "8.0"
        assert exit_code == 0

    def test_sub(self, monkeypatch, capsys):
        monkeypatch.setattr(sys, "argv", ["calc_3_12", "sub", "5", "3"])
        exit_code = main()
        captured = capsys.readouterr()
        assert captured.out.strip() == "2.0"
        assert exit_code == 0

    def test_mul(self, monkeypatch, capsys):
        monkeypatch.setattr(sys, "argv", ["calc_3_12", "mul", "2", "4"])
        exit_code = main()
        captured = capsys.readouterr()
        assert captured.out.strip() == "8.0"
        assert exit_code == 0

    def test_div(self, monkeypatch, capsys):
        monkeypatch.setattr(sys, "argv", ["calc_3_12", "div", "8", "4"])
        exit_code = main()
        captured = capsys.readouterr()
        assert captured.out.strip() == "2.0"
        assert exit_code == 0


class TestCLIErrorPaths:
    """Test error handling for divide-by-zero and invalid operations."""

    def test_divide_by_zero(self, monkeypatch, capsys):
        monkeypatch.setattr(sys, "argv", ["calc_3_12", "div", "5", "0"])
        exit_code = main()
        captured = capsys.readouterr()
        assert "Cannot divide by zero" in captured.err
        assert exit_code == 1

    def test_invalid_operation(self, monkeypatch):
        monkeypatch.setattr(sys, "argv", ["calc_3_12", "pow", "2", "3"])
        with pytest.raises(SystemExit) as exc_info:
            main()
        assert exc_info.value.code != 0


class TestCLISubprocess:
    """Verify the CLI works as a real subprocess invocation."""

    def test_add_subprocess(self):
        result = subprocess.run(
            [sys.executable, "-m", "calculator", "add", "5", "3"],
            capture_output=True,
            text=True,
        )
        assert result.stdout.strip() == "8.0"
        assert result.returncode == 0

    def test_divide_by_zero_subprocess(self):
        result = subprocess.run(
            [sys.executable, "-m", "calculator", "div", "5", "0"],
            capture_output=True,
            text=True,
        )
        assert "Cannot divide by zero" in result.stderr
        assert result.returncode == 1
