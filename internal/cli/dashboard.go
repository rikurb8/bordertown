package cli

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rikurb8/carnie/internal/dashboard"
	"github.com/spf13/cobra"
)

func newDashboardCommand() *cobra.Command {
	var refresh time.Duration
	var limit int

	cmd := &cobra.Command{
		Use:   "dashboard",
		Short: "Launch the Carnie dashboard",
		RunE: func(cmd *cobra.Command, args []string) error {
			model := dashboard.NewModel(refresh, limit)
			program := tea.NewProgram(model, tea.WithAltScreen())
			_, err := program.Run()
			return err
		},
	}

	cmd.Flags().DurationVar(&refresh, "refresh", 6*time.Second, "Auto-refresh interval (0 to disable)")
	cmd.Flags().IntVar(&limit, "limit", 200, "Max issues per column (0 for unlimited)")

	return cmd
}
