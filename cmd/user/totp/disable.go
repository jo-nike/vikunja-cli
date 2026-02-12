package totp

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newDisableCmd() *cobra.Command {
	var password string

	cmd := &cobra.Command{
		Use:   "disable",
		Short: "Disable TOTP",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{"password": password}
			if err := c.Post("/user/settings/totp/disable", body, nil); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "TOTP disabled"})
		},
	}
	cmd.Flags().StringVar(&password, "password", "", "Account password (required)")
	cmd.MarkFlagRequired("password")
	return cmd
}
