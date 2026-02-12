package members

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newAdminCmd() *cobra.Command {
	var teamID, userID int64
	var admin bool

	cmd := &cobra.Command{
		Use:   "admin",
		Short: "Toggle admin status of a team member",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{"admin": admin}
			if err := c.Update(fmt.Sprintf("/teams/%d/members/%d/admin", teamID, userID), body, nil); err != nil {
				output.Error(err)
			}
			output.Result(map[string]interface{}{"message": "admin status updated", "admin": admin})
		},
	}
	cmd.Flags().Int64Var(&teamID, "team-id", 0, "Team ID (required)")
	cmd.Flags().Int64Var(&userID, "user-id", 0, "User ID (required)")
	cmd.Flags().BoolVar(&admin, "admin", false, "Set admin status")
	cmd.MarkFlagRequired("team-id")
	cmd.MarkFlagRequired("user-id")
	return cmd
}
