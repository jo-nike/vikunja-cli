package user

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newPasswordCmd() *cobra.Command {
	var oldPassword, newPassword string

	cmd := &cobra.Command{
		Use:   "password",
		Short: "Change user password",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{
				"old_password": oldPassword,
				"new_password": newPassword,
			}
			if err := c.Post("/user/password", body, nil); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "password changed"})
		},
	}
	cmd.Flags().StringVar(&oldPassword, "old-password", "", "Current password (required)")
	cmd.Flags().StringVar(&newPassword, "new-password", "", "New password (required)")
	cmd.MarkFlagRequired("old-password")
	cmd.MarkFlagRequired("new-password")
	return cmd
}
