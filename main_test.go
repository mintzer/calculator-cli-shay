package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunHappyPaths(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantStdout string
		wantCode   int
	}{
		{"add integers", []string{"add", "5", "3"}, "8.0\n", 0},
		{"sub integers", []string{"sub", "10", "4"}, "6.0\n", 0},
		{"mul integers", []string{"mul", "6", "7"}, "42.0\n", 0},
		{"div exact", []string{"div", "20", "4"}, "5.0\n", 0},
		{"div fractional", []string{"div", "7", "3"}, "2.3333333333333335\n", 0},
		{"add fractional", []string{"add", "5.5", "3.5"}, "9.0\n", 0},
		{"mul by zero", []string{"mul", "0", "5"}, "0.0\n", 0},
		{"add negatives", []string{"add", "-3", "-7"}, "-10.0\n", 0},
		{"sub zeros", []string{"sub", "0", "0"}, "0.0\n", 0},
		{"div fractional inputs", []string{"div", "5.5", "2.5"}, "2.2\n", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := run(tt.args, &stdout, &stderr)
			if code != tt.wantCode {
				t.Errorf("run(%v) exit code = %d, want %d\nstderr: %s", tt.args, code, tt.wantCode, stderr.String())
			}
			if stdout.String() != tt.wantStdout {
				t.Errorf("run(%v) stdout = %q, want %q", tt.args, stdout.String(), tt.wantStdout)
			}
			if stderr.Len() > 0 {
				t.Errorf("run(%v) unexpected stderr: %s", tt.args, stderr.String())
			}
		})
	}
}

func TestRunDivideByZero(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"div", "5", "0"}, &stdout, &stderr)
	if code != 1 {
		t.Errorf("exit code = %d, want 1", code)
	}
	if stdout.Len() > 0 {
		t.Errorf("unexpected stdout: %s", stdout.String())
	}
	wantErr := "Error: Cannot divide by zero\n"
	if stderr.String() != wantErr {
		t.Errorf("stderr = %q, want %q", stderr.String(), wantErr)
	}
}

func TestRunInvalidOperation(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"foo", "5", "3"}, &stdout, &stderr)
	if code != 2 {
		t.Errorf("exit code = %d, want 2", code)
	}
	if stdout.Len() > 0 {
		t.Errorf("unexpected stdout: %s", stdout.String())
	}
	if !strings.Contains(stderr.String(), "invalid operation 'foo'") {
		t.Errorf("stderr should mention invalid operation, got: %s", stderr.String())
	}
}

func TestRunNonNumericArguments(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantMsg string
	}{
		{"non-numeric a", []string{"add", "abc", "5"}, "invalid float value for a: 'abc'"},
		{"non-numeric b", []string{"add", "5", "xyz"}, "invalid float value for b: 'xyz'"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := run(tt.args, &stdout, &stderr)
			if code != 2 {
				t.Errorf("exit code = %d, want 2", code)
			}
			if !strings.Contains(stderr.String(), tt.wantMsg) {
				t.Errorf("stderr = %q, want to contain %q", stderr.String(), tt.wantMsg)
			}
		})
	}
}

func TestRunMissingArguments(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantMsg  string
		wantCode int
	}{
		{"no args", []string{}, "operation, a, b", 2},
		{"only operation", []string{"add"}, "a, b", 2},
		{"operation and a", []string{"add", "5"}, "b", 2},
		{"too many args", []string{"add", "1", "2", "3"}, "too many arguments", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := run(tt.args, &stdout, &stderr)
			if code != tt.wantCode {
				t.Errorf("exit code = %d, want %d", code, tt.wantCode)
			}
			if !strings.Contains(stderr.String(), tt.wantMsg) {
				t.Errorf("stderr = %q, want to contain %q", stderr.String(), tt.wantMsg)
			}
		})
	}
}

func TestRunHelp(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"short flag", []string{"-h"}},
		{"long flag", []string{"--help"}},
		{"help with other args", []string{"add", "--help", "3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := run(tt.args, &stdout, &stderr)
			if code != 0 {
				t.Errorf("exit code = %d, want 0", code)
			}
			if !strings.Contains(stdout.String(), "Usage: calc-go") {
				t.Errorf("stdout should contain usage text, got: %s", stdout.String())
			}
			if stderr.Len() > 0 {
				t.Errorf("unexpected stderr: %s", stderr.String())
			}
		})
	}
}

func TestFormatFloat(t *testing.T) {
	tests := []struct {
		input float64
		want  string
	}{
		{8.0, "8.0"},
		{6.0, "6.0"},
		{42.0, "42.0"},
		{5.0, "5.0"},
		{0.0, "0.0"},
		{-10.0, "-10.0"},
		{2.3333333333333335, "2.3333333333333335"},
		{9.0, "9.0"},
		{2.2, "2.2"},
		{0.30000000000000004, "0.30000000000000004"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := formatFloat(tt.input)
			if got != tt.want {
				t.Errorf("formatFloat(%v) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
