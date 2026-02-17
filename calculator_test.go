package main

import (
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var calcBinary string

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "calc-test")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	calcBinary = filepath.Join(dir, "calc")
	cmd := exec.Command("go", "build", "-o", calcBinary, ".")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic("failed to build calc binary: " + err.Error())
	}

	os.Exit(m.Run())
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
		{"zero and positive", 0, 5, 5},
		{"positive and zero", 5, 0, 5},
		{"both zero", 0, 0, 0},
		{"fractional values", 1.5, 2.3, 3.8},
		{"large numbers", 1e10, 2e10, 3e10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Fatalf("Add(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 10, 4, 6},
		{"negative numbers", -2, -3, 1},
		{"mixed signs", -2, 3, -5},
		{"zero minus positive", 0, 5, -5},
		{"positive minus zero", 5, 0, 5},
		{"both zero", 0, 0, 0},
		{"fractional values", 5.5, 2.3, 3.2},
		{"result is negative", 3, 7, -4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Fatalf("Subtract(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 15},
		{"negative numbers", -2, -3, 6},
		{"mixed signs", -2, 3, -6},
		{"zero times positive", 0, 5, 0},
		{"positive times zero", 5, 0, 0},
		{"both zero", 0, 0, 0},
		{"fractional values", 2.5, 4, 10},
		{"one", 7, 1, 7},
		{"negative one", 7, -1, -7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Fatalf("Multiply(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a, b      float64
		expected  float64
		expectErr bool
	}{
		{"positive numbers", 20, 4, 5, false},
		{"negative numbers", -6, -3, 2, false},
		{"mixed signs", -6, 3, -2, false},
		{"fractional result", 7, 2, 3.5, false},
		{"fractional inputs", 2.5, 0.5, 5, false},
		{"zero numerator", 0, 5, 0, false},
		{"one", 9, 1, 9, false},
		{"divide by zero", 20, 0, 0, true},
		{"zero divided by zero", 0, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("Divide(%v, %v) expected error, got nil", tt.a, tt.b)
				}
				if err.Error() != "cannot divide by zero" {
					t.Fatalf("Divide(%v, %v) error = %q, want %q", tt.a, tt.b, err.Error(), "cannot divide by zero")
				}
			} else {
				if err != nil {
					t.Fatalf("Divide(%v, %v) unexpected error: %v", tt.a, tt.b, err)
				}
				if math.Abs(got-tt.expected) > 1e-9 {
					t.Fatalf("Divide(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
				}
			}
		})
	}
}

// --- CLI Integration Tests ---

func runCalc(args ...string) (stdout, stderr string, exitCode int) {
	cmd := exec.Command(calcBinary, args...)
	var outBuf, errBuf strings.Builder
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	err := cmd.Run()
	exitCode = 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			exitCode = -1
		}
	}
	return outBuf.String(), errBuf.String(), exitCode
}

func TestCLIAdd(t *testing.T) {
	stdout, _, exitCode := runCalc("add", "5", "3")
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}
	if strings.TrimSpace(stdout) != "8.0" {
		t.Fatalf("expected stdout %q, got %q", "8.0", strings.TrimSpace(stdout))
	}
}

func TestCLISub(t *testing.T) {
	stdout, _, exitCode := runCalc("sub", "10", "4")
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}
	if strings.TrimSpace(stdout) != "6.0" {
		t.Fatalf("expected stdout %q, got %q", "6.0", strings.TrimSpace(stdout))
	}
}

func TestCLIMul(t *testing.T) {
	stdout, _, exitCode := runCalc("mul", "2.5", "4")
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}
	if strings.TrimSpace(stdout) != "10.0" {
		t.Fatalf("expected stdout %q, got %q", "10.0", strings.TrimSpace(stdout))
	}
}

func TestCLIDiv(t *testing.T) {
	stdout, _, exitCode := runCalc("div", "20", "4")
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}
	if strings.TrimSpace(stdout) != "5.0" {
		t.Fatalf("expected stdout %q, got %q", "5.0", strings.TrimSpace(stdout))
	}
}

func TestCLIDivFractional(t *testing.T) {
	stdout, _, exitCode := runCalc("div", "7", "2")
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}
	if strings.TrimSpace(stdout) != "3.5" {
		t.Fatalf("expected stdout %q, got %q", "3.5", strings.TrimSpace(stdout))
	}
}

func TestCLIDivideByZero(t *testing.T) {
	_, stderr, exitCode := runCalc("div", "20", "0")
	if exitCode != 1 {
		t.Fatalf("expected exit code 1 for divide by zero, got %d", exitCode)
	}
	if !strings.Contains(stderr, "Cannot divide by zero") {
		t.Fatalf("expected stderr to contain %q, got %q", "Cannot divide by zero", stderr)
	}
}

func TestCLIUnknownOperation(t *testing.T) {
	_, stderr, exitCode := runCalc("foo", "1", "2")
	if exitCode != 2 {
		t.Fatalf("expected exit code 2 for unknown operation, got %d", exitCode)
	}
	if !strings.Contains(stderr, "invalid choice") {
		t.Fatalf("expected stderr to contain %q, got %q", "invalid choice", stderr)
	}
}

func TestCLIMissingArguments(t *testing.T) {
	_, stderr, exitCode := runCalc("add", "5")
	if exitCode != 2 {
		t.Fatalf("expected exit code 2 for missing arguments, got %d", exitCode)
	}
	if !strings.Contains(stderr, "arguments are required") {
		t.Fatalf("expected stderr to contain %q, got %q", "arguments are required", stderr)
	}
}

func TestCLIInvalidNumber(t *testing.T) {
	_, stderr, exitCode := runCalc("add", "abc", "3")
	if exitCode != 2 {
		t.Fatalf("expected exit code 2 for invalid number, got %d", exitCode)
	}
	if !strings.Contains(stderr, "invalid float value") {
		t.Fatalf("expected stderr to contain %q, got %q", "invalid float value", stderr)
	}
}

func TestCLIInvalidSecondNumber(t *testing.T) {
	_, stderr, exitCode := runCalc("add", "3", "xyz")
	if exitCode != 2 {
		t.Fatalf("expected exit code 2 for invalid number, got %d", exitCode)
	}
	if !strings.Contains(stderr, "invalid float value") {
		t.Fatalf("expected stderr to contain %q, got %q", "invalid float value", stderr)
	}
}

func TestCLINoArguments(t *testing.T) {
	_, stderr, exitCode := runCalc()
	if exitCode != 2 {
		t.Fatalf("expected exit code 2 for no arguments, got %d", exitCode)
	}
	if !strings.Contains(stderr, "arguments are required") {
		t.Fatalf("expected stderr to contain %q, got %q", "arguments are required", stderr)
	}
}

func TestCLIHelp(t *testing.T) {
	stdout, _, exitCode := runCalc("--help")
	if exitCode != 0 {
		t.Fatalf("expected exit code 0 for --help, got %d", exitCode)
	}
	if !strings.Contains(stdout, "usage: calc") {
		t.Fatalf("expected help output to contain %q, got %q", "usage: calc", stdout)
	}
}

func TestCLINegativeNumbers(t *testing.T) {
	stdout, _, exitCode := runCalc("add", "-2", "-3")
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}
	if strings.TrimSpace(stdout) != "-5.0" {
		t.Fatalf("expected stdout %q, got %q", "-5.0", strings.TrimSpace(stdout))
	}
}
