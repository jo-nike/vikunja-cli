package totp

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newEnrollCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "enroll",
		Short: "Enroll in TOTP (get secret and URL)",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var result interface{}
			if err := c.Post("/user/settings/totp/enroll", nil, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
}
