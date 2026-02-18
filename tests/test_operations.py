"""Unit tests for the calculator arithmetic functions."""

import pytest

from calculator import add, subtract, multiply, divide


class TestAdd:
    def test_positive_numbers(self):
        assert add(5, 3) == 8.0

    def test_negative_numbers(self):
        assert add(-2, -3) == -5.0

    def test_mixed_signs(self):
        assert add(-2, 5) == 3.0

    def test_zeros(self):
        assert add(0, 0) == 0.0

    def test_zero_and_positive(self):
        assert add(0, 7) == 7.0

    def test_floating_point(self):
        assert add(1.5, 2.5) == 4.0


class TestSubtract:
    def test_positive_numbers(self):
        assert subtract(5, 3) == 2.0

    def test_negative_numbers(self):
        assert subtract(-2, -3) == 1.0

    def test_mixed_signs(self):
        assert subtract(-2, 5) == -7.0

    def test_zeros(self):
        assert subtract(0, 0) == 0.0

    def test_zero_and_positive(self):
        assert subtract(0, 7) == -7.0

    def test_floating_point(self):
        assert subtract(5.5, 2.5) == 3.0


class TestMultiply:
    def test_positive_numbers(self):
        assert multiply(2, 4) == 8.0

    def test_negative_numbers(self):
        assert multiply(-2, -3) == 6.0

    def test_mixed_signs(self):
        assert multiply(-2, 5) == -10.0

    def test_zero(self):
        assert multiply(5, 0) == 0.0

    def test_both_zeros(self):
        assert multiply(0, 0) == 0.0

    def test_floating_point(self):
        assert multiply(1.5, 4.0) == 6.0


class TestDivide:
    def test_positive_numbers(self):
        assert divide(8, 4) == 2.0

    def test_negative_numbers(self):
        assert divide(-6, -3) == 2.0

    def test_mixed_signs(self):
        assert divide(-10, 5) == -2.0

    def test_zero_numerator(self):
        assert divide(0, 5) == 0.0

    def test_floating_point(self):
        assert divide(7.5, 2.5) == 3.0

    def test_divide_by_zero(self):
        with pytest.raises(ValueError, match="Cannot divide by zero"):
            divide(5, 0)

    def test_divide_zero_by_zero(self):
        with pytest.raises(ValueError, match="Cannot divide by zero"):
            divide(0, 0)
