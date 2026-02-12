package settings

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newEmailCmd() *cobra.Command {
	var email string

	cmd := &cobra.Command{
		Use:   "email",
		Short: "Update user email address",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{"new_email": email}
			if err := c.Post("/user/settings/email", body, nil); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "email update requested, check your inbox"})
		},
	}
	cmd.Flags().StringVar(&email, "email", "", "New email address (required)")
	cmd.MarkFlagRequired("email")
	return cmd
}
