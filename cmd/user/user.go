package user

import (
	"github.com/jo-nike/vikunja-cli/cmd/user/caldav_token"
	"github.com/jo-nike/vikunja-cli/cmd/user/deletion"
	"github.com/jo-nike/vikunja-cli/cmd/user/export"
	"github.com/jo-nike/vikunja-cli/cmd/user/settings"
	"github.com/jo-nike/vikunja-cli/cmd/user/totp"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Manage user account",
	}
	cmd.AddCommand(newGetCmd())
	cmd.AddCommand(newListCmd())
	cmd.AddCommand(newPasswordCmd())
	cmd.AddCommand(newTokenCmd())
	cmd.AddCommand(settings.NewCmd())
	cmd.AddCommand(export.NewCmd())
	cmd.AddCommand(deletion.NewCmd())
	cmd.AddCommand(totp.NewCmd())
	cmd.AddCommand(caldav_token.NewCmd())
	return cmd
}
