package members

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newAddCmd() *cobra.Command {
	var teamID, userID int64

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a member to a team",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{"user_id": userID}
			var result models.TeamMember
			if err := c.Create(fmt.Sprintf("/teams/%d/members", teamID), body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&teamID, "team-id", 0, "Team ID (required)")
	cmd.Flags().Int64Var(&userID, "user-id", 0, "User ID (required)")
	cmd.MarkFlagRequired("team-id")
	cmd.MarkFlagRequired("user-id")
	return cmd
}
