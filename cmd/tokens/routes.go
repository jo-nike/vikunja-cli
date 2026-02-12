package tokens

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newRoutesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "routes",
		Short: "List available API routes for token permissions",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var result interface{}
			if err := c.Get("/routes", &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
}
