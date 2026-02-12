package settings

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newTimezonesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "timezones",
		Short: "List available timezones",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var result interface{}
			if err := c.Get("/user/timezones", &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
}
