package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunHappyPath(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantOut    string
		wantErr    string
		wantCode   int
	}{
		{"add integers", []string{"add", "5", "3"}, "8.0\n", "", 0},
		{"sub integers", []string{"sub", "10", "4"}, "6.0\n", "", 0},
		{"mul integers", []string{"mul", "6", "7"}, "42.0\n", "", 0},
		{"div exact", []string{"div", "20", "4"}, "5.0\n", "", 0},
		{"div non-exact", []string{"div", "10", "3"}, "3.3333333333333335\n", "", 0},
		{"add floats", []string{"add", "1.5", "2.5"}, "4.0\n", "", 0},
		{"mul float and int", []string{"mul", "5.5", "2"}, "11.0\n", "", 0},
		{"sub same numbers", []string{"sub", "0", "0"}, "0.0\n", "", 0},
		{"add negative", []string{"add", "-5", "3"}, "-2.0\n", "", 0},
		{"mul negative", []string{"mul", "-3", "-4"}, "12.0\n", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := run(tt.args, &stdout, &stderr)
			if code != tt.wantCode {
				t.Errorf("run(%v) exit code = %d, want %d", tt.args, code, tt.wantCode)
			}
			if stdout.String() != tt.wantOut {
				t.Errorf("run(%v) stdout = %q, want %q", tt.args, stdout.String(), tt.wantOut)
			}
			if stderr.String() != tt.wantErr {
				t.Errorf("run(%v) stderr = %q, want %q", tt.args, stderr.String(), tt.wantErr)
			}
		})
	}
}

func TestRunDivideByZero(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"div", "10", "0"}, &stdout, &stderr)
	if code != 1 {
		t.Errorf("divide by zero exit code = %d, want 1", code)
	}
	if stdout.String() != "" {
		t.Errorf("divide by zero stdout = %q, want empty", stdout.String())
	}
	if !strings.Contains(stderr.String(), "Cannot divide by zero") {
		t.Errorf("divide by zero stderr = %q, want it to contain 'Cannot divide by zero'", stderr.String())
	}
}

func TestRunInvalidOperation(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"foo", "5", "3"}, &stdout, &stderr)
	if code != 2 {
		t.Errorf("invalid operation exit code = %d, want 2", code)
	}
	if !strings.Contains(stderr.String(), "invalid operation") {
		t.Errorf("invalid operation stderr = %q, want it to contain 'invalid operation'", stderr.String())
	}
}

func TestRunNonNumericArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"non-numeric a", []string{"add", "abc", "3"}},
		{"non-numeric b", []string{"add", "3", "xyz"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := run(tt.args, &stdout, &stderr)
			if code != 2 {
				t.Errorf("run(%v) exit code = %d, want 2", tt.args, code)
			}
			if !strings.Contains(stderr.String(), "invalid number") {
				t.Errorf("run(%v) stderr = %q, want it to contain 'invalid number'", tt.args, stderr.String())
			}
		})
	}
}

func TestRunMissingArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"no arguments", []string{}},
		{"only operation", []string{"add"}},
		{"missing b", []string{"add", "5"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := run(tt.args, &stdout, &stderr)
			if code != 2 {
				t.Errorf("run(%v) exit code = %d, want 2", tt.args, code)
			}
		})
	}
}

func TestRunExtraArgs(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"add", "5", "3", "7"}, &stdout, &stderr)
	if code != 2 {
		t.Errorf("extra args exit code = %d, want 2", code)
	}
	if !strings.Contains(stderr.String(), "too many arguments") {
		t.Errorf("extra args stderr = %q, want it to contain 'too many arguments'", stderr.String())
	}
}

func TestRunHelpFlag(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"short help", []string{"-h"}},
		{"long help", []string{"--help"}},
		{"help with other args", []string{"add", "-h", "3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := run(tt.args, &stdout, &stderr)
			if code != 0 {
				t.Errorf("run(%v) exit code = %d, want 0", tt.args, code)
			}
			if !strings.Contains(stdout.String(), "Usage: calc-go") {
				t.Errorf("run(%v) stdout = %q, want it to contain usage info", tt.args, stdout.String())
			}
		})
	}
}

func TestFormatResult(t *testing.T) {
	tests := []struct {
		input float64
		want  string
	}{
		{8.0, "8.0"},
		{0.0, "0.0"},
		{-2.0, "-2.0"},
		{42.0, "42.0"},
		{3.3333333333333335, "3.3333333333333335"},
		{11.0, "11.0"},
		{0.5, "0.5"},
		{-0.5, "-0.5"},
		{100.0, "100.0"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := formatResult(tt.input)
			if got != tt.want {
				t.Errorf("formatResult(%g) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
