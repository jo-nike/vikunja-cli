package notifications

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "notifications",
		Short: "Manage notifications",
	}
	cmd.AddCommand(newListCmd())
	cmd.AddCommand(newReadAllCmd())
	cmd.AddCommand(newReadCmd())
	return cmd
}
