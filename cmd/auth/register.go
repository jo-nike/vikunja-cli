package auth

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newRegisterCmd() *cobra.Command {
	var username, email, password string

	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register a new user account",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{
				"username": username,
				"email":    email,
				"password": password,
			}
			var result map[string]interface{}
			if err := c.Post("/register", body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().StringVar(&username, "username", "", "Username (required)")
	cmd.Flags().StringVar(&email, "email", "", "Email (required)")
	cmd.Flags().StringVar(&password, "password", "", "Password (required)")
	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("email")
	cmd.MarkFlagRequired("password")
	return cmd
}
