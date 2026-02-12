package projectteams

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newRemoveCmd() *cobra.Command {
	var projectID, teamID int64

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a team from a project",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			if err := c.Delete(fmt.Sprintf("/projects/%d/teams/%d", projectID, teamID)); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"status": "success"})
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().Int64Var(&teamID, "team-id", 0, "Team ID")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("team-id")

	return cmd
}
