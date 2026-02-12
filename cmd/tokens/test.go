package tokens

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newTestCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "Test the current API token",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var result map[string]interface{}
			if err := c.Get("/user", &result); err != nil {
				output.Error(err)
			}
			output.Result(map[string]interface{}{
				"status":   "ok",
				"username": result["username"],
				"id":       result["id"],
			})
		},
	}
}
