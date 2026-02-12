package migration

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migration",
		Short: "Manage data migrations from other services",
		Long:  "Import data from other services. Use --service to specify the service (todoist, trello, microsoft-todo, ticktick, vikunja-file).",
	}
	cmd.AddCommand(newAuthCmd())
	cmd.AddCommand(newMigrateCmd())
	cmd.AddCommand(newStatusCmd())
	return cmd
}
