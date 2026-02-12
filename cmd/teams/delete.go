package teams

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
		Short: "Delete a team",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			if err := c.Delete(fmt.Sprintf("/teams/%d", id)); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "team deleted"})
		},
	}
	cmd.Flags().Int64Var(&id, "id", 0, "Team ID (required)")
	cmd.MarkFlagRequired("id")
	return cmd
}
