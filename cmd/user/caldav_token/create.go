package caldav_token

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a CalDAV token",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var result models.CaldavToken
			if err := c.Create("/user/settings/token/caldav", nil, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
}
