package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantStdout string
		wantStderr string
		wantCode   int
	}{
		// Happy path — all four operations
		{
			name:       "add positive integers",
			args:       []string{"add", "5", "3"},
			wantStdout: "8.0\n",
			wantCode:   0,
		},
		{
			name:       "sub positive integers",
			args:       []string{"sub", "10", "4"},
			wantStdout: "6.0\n",
			wantCode:   0,
		},
		{
			name:       "mul positive integers",
			args:       []string{"mul", "6", "7"},
			wantStdout: "42.0\n",
			wantCode:   0,
		},
		{
			name:       "div even division",
			args:       []string{"div", "20", "4"},
			wantStdout: "5.0\n",
			wantCode:   0,
		},
		{
			name:       "div decimal result",
			args:       []string{"div", "10", "3"},
			wantStdout: "3.3333333333333335\n",
			wantCode:   0,
		},

		// Decimal and negative inputs
		{
			name:       "add negative number",
			args:       []string{"add", "-5", "3"},
			wantStdout: "-2.0\n",
			wantCode:   0,
		},
		{
			name:       "div decimal inputs",
			args:       []string{"div", "7.5", "2.5"},
			wantStdout: "3.0\n",
			wantCode:   0,
		},
		{
			name:       "mul with zero",
			args:       []string{"mul", "5", "0"},
			wantStdout: "0.0\n",
			wantCode:   0,
		},

		// Divide by zero
		{
			name:       "divide by zero",
			args:       []string{"div", "10", "0"},
			wantStderr: "Error: Cannot divide by zero\n",
			wantCode:   1,
		},

		// Invalid operation
		{
			name:       "invalid operation",
			args:       []string{"foo", "5", "3"},
			wantStderr: "calc: error: argument operation: invalid choice: 'foo' (choose from add, sub, mul, div)",
			wantCode:   2,
		},

		// Non-numeric arguments
		{
			name:       "non-numeric first argument",
			args:       []string{"add", "abc", "3"},
			wantStderr: "calc: error: argument a: invalid float value: 'abc'",
			wantCode:   2,
		},
		{
			name:       "non-numeric second argument",
			args:       []string{"add", "5", "xyz"},
			wantStderr: "calc: error: argument b: invalid float value: 'xyz'",
			wantCode:   2,
		},

		// Missing arguments
		{
			name:       "no arguments",
			args:       []string{},
			wantStderr: "calc: error: the following arguments are required: operation, a, b",
			wantCode:   2,
		},
		{
			name:       "only operation",
			args:       []string{"add"},
			wantStderr: "calc: error: the following arguments are required: a, b",
			wantCode:   2,
		},
		{
			name:       "missing second operand",
			args:       []string{"add", "5"},
			wantStderr: "calc: error: the following arguments are required: b",
			wantCode:   2,
		},

		// Extra arguments
		{
			name:       "extra arguments",
			args:       []string{"add", "5", "3", "extra"},
			wantStderr: "calc: error: unrecognized arguments: extra",
			wantCode:   2,
		},

		// Help flags
		{
			name:       "help short flag",
			args:       []string{"-h"},
			wantStdout: "usage: calc [-h] {add,sub,mul,div} a b",
			wantCode:   0,
		},
		{
			name:       "help long flag",
			args:       []string{"--help"},
			wantStdout: "usage: calc [-h] {add,sub,mul,div} a b",
			wantCode:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := run(tt.args, &stdout, &stderr)

			if code != tt.wantCode {
				t.Errorf("run(%v) exit code = %d, want %d", tt.args, code, tt.wantCode)
			}

			if tt.wantStdout != "" && !strings.Contains(stdout.String(), tt.wantStdout) {
				t.Errorf("run(%v) stdout = %q, want to contain %q", tt.args, stdout.String(), tt.wantStdout)
			}

			if tt.wantStderr != "" && !strings.Contains(stderr.String(), tt.wantStderr) {
				t.Errorf("run(%v) stderr = %q, want to contain %q", tt.args, stderr.String(), tt.wantStderr)
			}
		})
	}
}

func TestFormatResult(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{"integer value", 8.0, "8.0"},
		{"negative integer", -2.0, "-2.0"},
		{"zero", 0.0, "0.0"},
		{"decimal value", 3.14, "3.14"},
		{"repeating decimal", 10.0 / 3.0, "3.3333333333333335"},
		{"large integer", 1e10, "10000000000.0"},
		{"scientific notation", 1e16, "1e+16"},
		{"small decimal", 0.0001, "0.0001"},
		{"very small number", 1e-5, "1e-05"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatResult(tt.input)
			if result != tt.expected {
				t.Errorf("formatResult(%v) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
