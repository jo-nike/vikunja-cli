package auth

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authentication commands",
	}
	cmd.AddCommand(newLoginCmd())
	cmd.AddCommand(newRegisterCmd())
	return cmd
}
