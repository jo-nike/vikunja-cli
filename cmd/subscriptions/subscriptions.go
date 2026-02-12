package subscriptions

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscriptions",
		Short: "Manage entity subscriptions",
	}
	cmd.AddCommand(newCreateCmd())
	cmd.AddCommand(newDeleteCmd())
	return cmd
}
