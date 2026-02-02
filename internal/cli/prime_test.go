package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrimeOperatorRole(t *testing.T) {
	root := NewRootCommand()
	output := &bytes.Buffer{}
	root.SetOut(output)
	root.SetErr(output)
	root.SetArgs([]string{"prime", "operator"})

	if err := root.Execute(); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !strings.Contains(output.String(), "Operator") {
		t.Error("expected output to contain 'Operator'")
	}
}

func TestPrimeCarnieRole(t *testing.T) {
	root := NewRootCommand()
	output := &bytes.Buffer{}
	root.SetOut(output)
	root.SetErr(output)
	root.SetArgs([]string{"prime", "carnie"})

	if err := root.Execute(); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !strings.Contains(output.String(), "Carnie") {
		t.Error("expected output to contain 'Carnie'")
	}
}

func TestPrimeInvalidRole(t *testing.T) {
	root := NewRootCommand()
	output := &bytes.Buffer{}
	root.SetOut(output)
	root.SetErr(output)
	root.SetArgs([]string{"prime", "invalid"})

	err := root.Execute()
	if err == nil {
		t.Fatal("expected error for invalid role")
	}

	if !strings.Contains(err.Error(), "invalid role") {
		t.Errorf("expected 'invalid role' error, got: %v", err)
	}
	if !strings.Contains(err.Error(), "operator") {
		t.Error("expected error to list valid roles")
	}
}

func TestPrimeNoArgs(t *testing.T) {
	root := NewRootCommand()
	output := &bytes.Buffer{}
	root.SetOut(output)
	root.SetErr(output)
	root.SetArgs([]string{"prime"})

	err := root.Execute()
	if err == nil {
		t.Fatal("expected error when no role provided")
	}
}
