package user

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newTokenCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "token",
		Short: "Get a new JWT token for the current user",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var result map[string]interface{}
			if err := c.Post("/user/token", nil, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
}
