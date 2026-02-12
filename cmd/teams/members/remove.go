package members

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newRemoveCmd() *cobra.Command {
	var teamID int64
	var username string

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a member from a team",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			if err := c.Delete(fmt.Sprintf("/teams/%d/members/%s", teamID, username)); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "member removed"})
		},
	}
	cmd.Flags().Int64Var(&teamID, "team-id", 0, "Team ID (required)")
	cmd.Flags().StringVar(&username, "username", "", "Username to remove (required)")
	cmd.MarkFlagRequired("team-id")
	cmd.MarkFlagRequired("username")
	return cmd
}
