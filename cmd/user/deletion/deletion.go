package deletion

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deletion",
		Short: "Manage account deletion",
	}
	cmd.AddCommand(newRequestCmd())
	cmd.AddCommand(newConfirmCmd())
	cmd.AddCommand(newCancelCmd())
	return cmd
}
