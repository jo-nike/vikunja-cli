package settings

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "settings",
		Short: "Manage user settings",
	}
	cmd.AddCommand(newGeneralCmd())
	cmd.AddCommand(newEmailCmd())
	cmd.AddCommand(newAvatarCmd())
	cmd.AddCommand(newTimezonesCmd())
	return cmd
}
