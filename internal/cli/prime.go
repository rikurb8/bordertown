package cli

import (
	"fmt"
	"strings"

	"github.com/rikurb8/carnie/internal/prime"
	"github.com/spf13/cobra"
)

func newPrimeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prime <role>",
		Short: "Print role-specific workflow context",
		Long:  "Outputs role-specific context in markdown format for AI agent workflow recovery.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			roleName := args[0]

			if !prime.IsValidRole(roleName) {
				validRoles := make([]string, len(prime.ValidRoles()))
				for i, r := range prime.ValidRoles() {
					validRoles[i] = string(r)
				}
				return fmt.Errorf("invalid role %q, valid roles: %s", roleName, strings.Join(validRoles, ", "))
			}

			content, err := prime.LoadPrompt(prime.Role(roleName))
			if err != nil {
				return err
			}

			fmt.Fprint(cmd.OutOrStdout(), content)
			return nil
		},
	}

	return cmd
}
