package totp

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "totp",
		Short: "Manage TOTP two-factor authentication",
	}
	cmd.AddCommand(newStatusCmd())
	cmd.AddCommand(newEnrollCmd())
	cmd.AddCommand(newEnableCmd())
	cmd.AddCommand(newDisableCmd())
	cmd.AddCommand(newQRCodeCmd())
	return cmd
}
