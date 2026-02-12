package auth

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newLoginCmd() *cobra.Command {
	var username, password, totpPasscode string
	var longToken bool

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login and get a JWT token",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{
				"username": username,
				"password": password,
			}
			if cmd.Flags().Changed("totp-passcode") {
				body["totp_passcode"] = totpPasscode
			}
			if longToken {
				body["long_token"] = true
			}
			var result map[string]interface{}
			if err := c.Post("/login", body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().StringVar(&username, "username", "", "Username or email (required)")
	cmd.Flags().StringVar(&password, "password", "", "Password (required)")
	cmd.Flags().StringVar(&totpPasscode, "totp-passcode", "", "TOTP passcode")
	cmd.Flags().BoolVar(&longToken, "long-token", false, "Request a long-lived token")
	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")
	return cmd
}
