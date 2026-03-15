"""Unit tests for the calculator CLI."""

import io
import sys
import unittest
from unittest.mock import patch

from calculator import add, divide, main, multiply, subtract


class TestAdd(unittest.TestCase):
    """Tests for the add function."""

    def test_positive_numbers(self):
        self.assertEqual(add(2, 3), 5)

    def test_negative_numbers(self):
        self.assertEqual(add(-1, -2), -3)

    def test_mixed_signs(self):
        self.assertEqual(add(-1, 3), 2)

    def test_zeros(self):
        self.assertEqual(add(0, 0), 0)

    def test_floats(self):
        self.assertAlmostEqual(add(1.5, 2.3), 3.8)


class TestSubtract(unittest.TestCase):
    """Tests for the subtract function."""

    def test_positive_numbers(self):
        self.assertEqual(subtract(5, 3), 2)

    def test_negative_result(self):
        self.assertEqual(subtract(3, 5), -2)

    def test_negative_numbers(self):
        self.assertEqual(subtract(-1, -2), 1)

    def test_zeros(self):
        self.assertEqual(subtract(0, 0), 0)

    def test_floats(self):
        self.assertAlmostEqual(subtract(5.5, 2.3), 3.2)


class TestMultiply(unittest.TestCase):
    """Tests for the multiply function."""

    def test_positive_numbers(self):
        self.assertEqual(multiply(2, 3), 6)

    def test_by_zero(self):
        self.assertEqual(multiply(5, 0), 0)

    def test_negative_numbers(self):
        self.assertEqual(multiply(-2, -3), 6)

    def test_mixed_signs(self):
        self.assertEqual(multiply(-2, 3), -6)

    def test_floats(self):
        self.assertAlmostEqual(multiply(1.5, 2.0), 3.0)


class TestDivide(unittest.TestCase):
    """Tests for the divide function."""

    def test_positive_numbers(self):
        self.assertEqual(divide(6, 3), 2.0)

    def test_negative_numbers(self):
        self.assertEqual(divide(-6, -3), 2.0)

    def test_mixed_signs(self):
        self.assertEqual(divide(-6, 3), -2.0)

    def test_floats(self):
        self.assertAlmostEqual(divide(7.0, 2.0), 3.5)

    def test_divide_by_zero(self):
        with self.assertRaises(ValueError) as ctx:
            divide(1, 0)
        self.assertIn("Cannot divide by zero", str(ctx.exception))


class TestMain(unittest.TestCase):
    """Tests for the CLI main() entry point."""

    def test_add_operation(self):
        with patch("sys.argv", ["calc", "add", "2", "3"]):
            with patch("sys.stdout", new_callable=io.StringIO) as mock_out:
                result = main()
        self.assertEqual(result, 0)
        self.assertEqual(mock_out.getvalue().strip(), "5.0")

    def test_subtract_operation(self):
        with patch("sys.argv", ["calc", "sub", "10", "4"]):
            with patch("sys.stdout", new_callable=io.StringIO) as mock_out:
                result = main()
        self.assertEqual(result, 0)
        self.assertEqual(mock_out.getvalue().strip(), "6.0")

    def test_multiply_operation(self):
        with patch("sys.argv", ["calc", "mul", "3", "4"]):
            with patch("sys.stdout", new_callable=io.StringIO) as mock_out:
                result = main()
        self.assertEqual(result, 0)
        self.assertEqual(mock_out.getvalue().strip(), "12.0")

    def test_divide_operation(self):
        with patch("sys.argv", ["calc", "div", "10", "4"]):
            with patch("sys.stdout", new_callable=io.StringIO) as mock_out:
                result = main()
        self.assertEqual(result, 0)
        self.assertEqual(mock_out.getvalue().strip(), "2.5")

    def test_divide_by_zero_error(self):
        with patch("sys.argv", ["calc", "div", "1", "0"]):
            with patch("sys.stderr", new_callable=io.StringIO) as mock_err:
                result = main()
        self.assertEqual(result, 1)
        self.assertIn("Cannot divide by zero", mock_err.getvalue())


if __name__ == "__main__":
    unittest.main()
