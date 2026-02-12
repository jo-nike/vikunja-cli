package filters

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	var id int64

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a saved filter",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			if err := c.Delete(fmt.Sprintf("/filters/%d", id)); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "filter deleted"})
		},
	}
	cmd.Flags().Int64Var(&id, "id", 0, "Filter ID (required)")
	cmd.MarkFlagRequired("id")
	return cmd
}
