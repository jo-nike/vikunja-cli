package totp

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newEnableCmd() *cobra.Command {
	var passcode string

	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enable TOTP with a passcode",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{"passcode": passcode}
			var result interface{}
			if err := c.Post("/user/settings/totp/enable", body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().StringVar(&passcode, "passcode", "", "TOTP passcode (required)")
	cmd.MarkFlagRequired("passcode")
	return cmd
}
