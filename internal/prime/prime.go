package prime

import (
	"fmt"

	"github.com/rikurb8/carnie/internal/templates"
)

// Role represents a supported prime role.
type Role string

const (
	RoleOperator Role = "operator"
	RoleCarnie   Role = "carnie"
)

var validRoles = []Role{RoleOperator, RoleCarnie}

// ValidRoles returns all valid role names.
func ValidRoles() []Role {
	return validRoles
}

// IsValidRole checks if the given string is a valid role.
func IsValidRole(r string) bool {
	for _, role := range validRoles {
		if string(role) == r {
			return true
		}
	}
	return false
}

// LoadPrompt loads the embedded template for the given role.
func LoadPrompt(role Role) (string, error) {
	filename := string(role) + ".md"
	content, err := templates.Load(filename)
	if err != nil {
		return "", fmt.Errorf("load template for role %q: %w", role, err)
	}
	return content, nil
}
