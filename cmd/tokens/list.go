package tokens

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all API tokens",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var result []models.APIToken
			if err := c.Get("/tokens", &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
}
