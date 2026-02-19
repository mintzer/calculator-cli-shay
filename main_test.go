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
		wantStdout string
		wantExit   int
	}{
		{"add integers", []string{"add", "5", "3"}, "8.0\n", 0},
		{"sub integers", []string{"sub", "10", "4"}, "6.0\n", 0},
		{"mul integers", []string{"mul", "6", "7"}, "42.0\n", 0},
		{"div integers", []string{"div", "20", "4"}, "5.0\n", 0},
		{"div fractional result", []string{"div", "7", "3"}, "2.3333333333333335\n", 0},
		{"add zeros", []string{"add", "0", "0"}, "0.0\n", 0},
		{"mul with float", []string{"mul", "3.5", "2"}, "7.0\n", 0},
		{"sub negative result", []string{"sub", "3", "5"}, "-2.0\n", 0},
		{"mul negative numbers", []string{"mul", "-3", "-4"}, "12.0\n", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			exitCode := run(tt.args, &stdout, &stderr)
			if exitCode != tt.wantExit {
				t.Errorf("run(%v) exit code = %d, want %d", tt.args, exitCode, tt.wantExit)
			}
			if stdout.String() != tt.wantStdout {
				t.Errorf("run(%v) stdout = %q, want %q", tt.args, stdout.String(), tt.wantStdout)
			}
			if stderr.Len() > 0 {
				t.Errorf("run(%v) unexpected stderr output: %q", tt.args, stderr.String())
			}
		})
	}
}

func TestRunDivideByZero(t *testing.T) {
	var stdout, stderr bytes.Buffer
	exitCode := run([]string{"div", "5", "0"}, &stdout, &stderr)
	if exitCode != 1 {
		t.Errorf("div by zero: exit code = %d, want 1", exitCode)
	}
	if stdout.Len() > 0 {
		t.Errorf("div by zero: unexpected stdout: %q", stdout.String())
	}
	wantErr := "Error: Cannot divide by zero"
	if !strings.Contains(stderr.String(), wantErr) {
		t.Errorf("div by zero: stderr = %q, want it to contain %q", stderr.String(), wantErr)
	}
}

func TestRunInvalidOperation(t *testing.T) {
	var stdout, stderr bytes.Buffer
	exitCode := run([]string{"foo", "5", "3"}, &stdout, &stderr)
	if exitCode != 2 {
		t.Errorf("invalid op: exit code = %d, want 2", exitCode)
	}
	if stdout.Len() > 0 {
		t.Errorf("invalid op: unexpected stdout: %q", stdout.String())
	}
	if !strings.Contains(stderr.String(), "invalid operation") {
		t.Errorf("invalid op: stderr = %q, want it to mention invalid operation", stderr.String())
	}
}

func TestRunNonNumericArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"non-numeric a", []string{"add", "abc", "3"}, "invalid float value for a"},
		{"non-numeric b", []string{"add", "3", "xyz"}, "invalid float value for b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			exitCode := run(tt.args, &stdout, &stderr)
			if exitCode != 2 {
				t.Errorf("run(%v) exit code = %d, want 2", tt.args, exitCode)
			}
			if !strings.Contains(stderr.String(), tt.want) {
				t.Errorf("run(%v) stderr = %q, want it to contain %q", tt.args, stderr.String(), tt.want)
			}
		})
	}
}

func TestRunMissingArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"no args", []string{}},
		{"only operation", []string{"add"}},
		{"only two args", []string{"add", "5"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			exitCode := run(tt.args, &stdout, &stderr)
			if exitCode != 2 {
				t.Errorf("run(%v) exit code = %d, want 2", tt.args, exitCode)
			}
			if !strings.Contains(stderr.String(), "Usage:") {
				t.Errorf("run(%v) stderr = %q, want it to contain usage info", tt.args, stderr.String())
			}
		})
	}
}

func TestRunExtraArgs(t *testing.T) {
	var stdout, stderr bytes.Buffer
	exitCode := run([]string{"add", "5", "3", "7"}, &stdout, &stderr)
	if exitCode != 2 {
		t.Errorf("extra args: exit code = %d, want 2", exitCode)
	}
}

func TestRunHelp(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"short flag", []string{"-h"}},
		{"long flag", []string{"--help"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			exitCode := run(tt.args, &stdout, &stderr)
			if exitCode != 0 {
				t.Errorf("run(%v) exit code = %d, want 0", tt.args, exitCode)
			}
			if !strings.Contains(stdout.String(), "Usage:") {
				t.Errorf("run(%v) stdout = %q, want it to contain usage info", tt.args, stdout.String())
			}
			if !strings.Contains(stdout.String(), "add") || !strings.Contains(stdout.String(), "div") {
				t.Errorf("run(%v) stdout should list operations", tt.args)
			}
			if stderr.Len() > 0 {
				t.Errorf("run(%v) unexpected stderr: %q", tt.args, stderr.String())
			}
		})
	}
}

func TestFormatFloat(t *testing.T) {
	tests := []struct {
		input float64
		want  string
	}{
		{8, "8.0"},
		{0, "0.0"},
		{-5, "-5.0"},
		{42, "42.0"},
		{8.5, "8.5"},
		{2.3333333333333335, "2.3333333333333335"},
		{0.1, "0.1"},
		{-3.14, "-3.14"},
		{1e20, "1e+20"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := formatFloat(tt.input)
			if got != tt.want {
				t.Errorf("formatFloat(%g) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
