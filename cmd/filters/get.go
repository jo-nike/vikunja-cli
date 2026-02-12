package filters

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	var id int64

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a saved filter by ID",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var result models.SavedFilter
			if err := c.Get(fmt.Sprintf("/filters/%d", id), &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&id, "id", 0, "Filter ID (required)")
	cmd.MarkFlagRequired("id")
	return cmd
}
