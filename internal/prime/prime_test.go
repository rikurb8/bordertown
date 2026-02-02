package prime

import (
	"strings"
	"testing"
)

func TestValidRoles(t *testing.T) {
	roles := ValidRoles()
	if len(roles) != 2 {
		t.Errorf("expected 2 roles, got %d", len(roles))
	}

	hasOperator := false
	hasCarnie := false
	for _, r := range roles {
		if r == RoleOperator {
			hasOperator = true
		}
		if r == RoleCarnie {
			hasCarnie = true
		}
	}

	if !hasOperator {
		t.Error("expected operator role in ValidRoles()")
	}
	if !hasCarnie {
		t.Error("expected carnie role in ValidRoles()")
	}
}

func TestIsValidRole(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"operator", true},
		{"carnie", true},
		{"invalid", false},
		{"", false},
		{"OPERATOR", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := IsValidRole(tt.input)
			if got != tt.want {
				t.Errorf("IsValidRole(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestLoadPromptOperator(t *testing.T) {
	content, err := LoadPrompt(RoleOperator)
	if err != nil {
		t.Fatalf("LoadPrompt(operator) error: %v", err)
	}

	if !strings.Contains(content, "Operator") {
		t.Error("expected operator template to contain 'Operator'")
	}
	if !strings.Contains(content, "bd") {
		t.Error("expected operator template to mention 'bd' commands")
	}
}

func TestLoadPromptCarnie(t *testing.T) {
	content, err := LoadPrompt(RoleCarnie)
	if err != nil {
		t.Fatalf("LoadPrompt(carnie) error: %v", err)
	}

	if !strings.Contains(content, "Carnie") {
		t.Error("expected carnie template to contain 'Carnie'")
	}
	if !strings.Contains(content, "implementation") {
		t.Error("expected carnie template to mention 'implementation'")
	}
}
