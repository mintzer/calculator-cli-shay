package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// binaryPath holds the path to the compiled calc binary used by tests.
var binaryPath string

// TestMain builds the binary once before all tests in this package run.
func TestMain(m *testing.M) {
	// Build binary into a temp directory.
	dir, err := os.MkdirTemp("", "calc-test-*")
	if err != nil {
		panic("failed to create temp dir: " + err.Error())
	}
	defer os.RemoveAll(dir)

	binaryName := "calc"
	if runtime.GOOS == "windows" {
		binaryName = "calc.exe"
	}
	binaryPath = filepath.Join(dir, binaryName)

	cmd := exec.Command("go", "build", "-o", binaryPath, ".")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic("failed to build binary: " + err.Error())
	}

	os.Exit(m.Run())
}

// runCalc executes the calc binary with the given arguments and returns
// stdout, stderr, and the exit code.
func runCalc(args ...string) (stdout, stderr string, exitCode int) {
	cmd := exec.Command(binaryPath, args...)
	var outBuf, errBuf bytes.Buffer
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

func TestSuccessfulAdd(t *testing.T) {
	stdout, stderr, exitCode := runCalc("add", "5", "3")
	if exitCode != 0 {
		t.Errorf("exit code = %d, want 0", exitCode)
	}
	if strings.TrimSpace(stdout) != "8.0" {
		t.Errorf("stdout = %q, want %q", stdout, "8.0\n")
	}
	if stderr != "" {
		t.Errorf("stderr = %q, want empty", stderr)
	}
}

func TestSuccessfulSub(t *testing.T) {
	stdout, stderr, exitCode := runCalc("sub", "10", "3")
	if exitCode != 0 {
		t.Errorf("exit code = %d, want 0", exitCode)
	}
	if strings.TrimSpace(stdout) != "7.0" {
		t.Errorf("stdout = %q, want %q", stdout, "7.0\n")
	}
	if stderr != "" {
		t.Errorf("stderr = %q, want empty", stderr)
	}
}

func TestSuccessfulMul(t *testing.T) {
	stdout, stderr, exitCode := runCalc("mul", "2", "4")
	if exitCode != 0 {
		t.Errorf("exit code = %d, want 0", exitCode)
	}
	if strings.TrimSpace(stdout) != "8.0" {
		t.Errorf("stdout = %q, want %q", stdout, "8.0\n")
	}
	if stderr != "" {
		t.Errorf("stderr = %q, want empty", stderr)
	}
}

func TestSuccessfulDiv(t *testing.T) {
	stdout, stderr, exitCode := runCalc("div", "20", "4")
	if exitCode != 0 {
		t.Errorf("exit code = %d, want 0", exitCode)
	}
	if strings.TrimSpace(stdout) != "5.0" {
		t.Errorf("stdout = %q, want %q", stdout, "5.0\n")
	}
	if stderr != "" {
		t.Errorf("stderr = %q, want empty", stderr)
	}
}

func TestSuccessfulFractional(t *testing.T) {
	stdout, _, exitCode := runCalc("add", "2.5", "3.1")
	if exitCode != 0 {
		t.Errorf("exit code = %d, want 0", exitCode)
	}
	if strings.TrimSpace(stdout) != "5.6" {
		t.Errorf("stdout = %q, want %q", stdout, "5.6\n")
	}
}

func TestSuccessfulNegativeNumbers(t *testing.T) {
	stdout, _, exitCode := runCalc("add", "-1", "1")
	if exitCode != 0 {
		t.Errorf("exit code = %d, want 0", exitCode)
	}
	if strings.TrimSpace(stdout) != "0.0" {
		t.Errorf("stdout = %q, want %q", stdout, "0.0\n")
	}
}

func TestDivisionByZero(t *testing.T) {
	stdout, stderr, exitCode := runCalc("div", "20", "0")
	if exitCode != 1 {
		t.Errorf("exit code = %d, want 1", exitCode)
	}
	if !strings.Contains(stderr, "Cannot divide by zero") {
		t.Errorf("stderr = %q, want it to contain %q", stderr, "Cannot divide by zero")
	}
	if stdout != "" {
		t.Errorf("stdout = %q, want empty", stdout)
	}
}

func TestInvalidOperation(t *testing.T) {
	stdout, stderr, exitCode := runCalc("foo", "1", "2")
	if exitCode == 0 {
		t.Error("exit code = 0, want non-zero")
	}
	if !strings.Contains(stderr, "invalid operation") {
		t.Errorf("stderr = %q, want it to contain %q", stderr, "invalid operation")
	}
	if !strings.Contains(stderr, "Usage:") {
		t.Errorf("stderr = %q, want it to contain usage text", stderr)
	}
	if stdout != "" {
		t.Errorf("stdout = %q, want empty", stdout)
	}
}

func TestMissingArguments(t *testing.T) {
	stdout, stderr, exitCode := runCalc("add", "1")
	if exitCode == 0 {
		t.Error("exit code = 0, want non-zero")
	}
	if !strings.Contains(stderr, "Usage:") {
		t.Errorf("stderr = %q, want it to contain usage text", stderr)
	}
	if stdout != "" {
		t.Errorf("stdout = %q, want empty", stdout)
	}
}

func TestNoArguments(t *testing.T) {
	stdout, stderr, exitCode := runCalc()
	if exitCode == 0 {
		t.Error("exit code = 0, want non-zero")
	}
	if !strings.Contains(stderr, "Usage:") {
		t.Errorf("stderr = %q, want it to contain usage text", stderr)
	}
	if stdout != "" {
		t.Errorf("stdout = %q, want empty", stdout)
	}
}

func TestNonNumericOperands(t *testing.T) {
	stdout, stderr, exitCode := runCalc("add", "one", "2")
	if exitCode == 0 {
		t.Error("exit code = 0, want non-zero")
	}
	if !strings.Contains(stderr, "invalid number") {
		t.Errorf("stderr = %q, want it to contain %q", stderr, "invalid number")
	}
	if stdout != "" {
		t.Errorf("stdout = %q, want empty", stdout)
	}
}

func TestNonNumericSecondOperand(t *testing.T) {
	stdout, stderr, exitCode := runCalc("add", "1", "two")
	if exitCode == 0 {
		t.Error("exit code = 0, want non-zero")
	}
	if !strings.Contains(stderr, "invalid number") {
		t.Errorf("stderr = %q, want it to contain %q", stderr, "invalid number")
	}
	if stdout != "" {
		t.Errorf("stdout = %q, want empty", stdout)
	}
}
