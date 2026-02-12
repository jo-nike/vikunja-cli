package cmd

import (
	"github.com/jo-nike/vikunja-cli/cmd/auth"
	"github.com/jo-nike/vikunja-cli/cmd/filters"
	"github.com/jo-nike/vikunja-cli/cmd/labels"
	"github.com/jo-nike/vikunja-cli/cmd/migration"
	"github.com/jo-nike/vikunja-cli/cmd/notifications"
	"github.com/jo-nike/vikunja-cli/cmd/projects"
	"github.com/jo-nike/vikunja-cli/cmd/reactions"
	"github.com/jo-nike/vikunja-cli/cmd/subscriptions"
	"github.com/jo-nike/vikunja-cli/cmd/system"
	"github.com/jo-nike/vikunja-cli/cmd/tasks"
	"github.com/jo-nike/vikunja-cli/cmd/teams"
	"github.com/jo-nike/vikunja-cli/cmd/tokens"
	"github.com/jo-nike/vikunja-cli/cmd/user"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "vikunja",
	Short:         "CLI for the Vikunja API",
	Long:          "A command-line interface for interacting with the Vikunja project management API. All output is JSON.",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		output.Error(err)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(system.NewCmd())
	rootCmd.AddCommand(projects.NewCmd())
	rootCmd.AddCommand(tasks.NewCmd())
	rootCmd.AddCommand(labels.NewCmd())
	rootCmd.AddCommand(teams.NewCmd())
	rootCmd.AddCommand(filters.NewCmd())
	rootCmd.AddCommand(auth.NewCmd())
	rootCmd.AddCommand(tokens.NewCmd())
	rootCmd.AddCommand(user.NewCmd())
	rootCmd.AddCommand(notifications.NewCmd())
	rootCmd.AddCommand(subscriptions.NewCmd())
	rootCmd.AddCommand(reactions.NewCmd())
	rootCmd.AddCommand(migration.NewCmd())
}
