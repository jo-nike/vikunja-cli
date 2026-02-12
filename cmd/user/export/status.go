package export

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Check export status",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var result interface{}
			if err := c.Get("/user/export", &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
}
