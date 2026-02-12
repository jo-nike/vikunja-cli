package relations

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relations",
		Short: "Manage task relations",
	}
	cmd.AddCommand(newCreateCmd())
	cmd.AddCommand(newDeleteCmd())
	return cmd
}
