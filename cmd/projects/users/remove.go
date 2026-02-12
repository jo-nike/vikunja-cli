package users

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newRemoveCmd() *cobra.Command {
	var projectID, userID int64

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a user from a project",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			if err := c.Delete(fmt.Sprintf("/projects/%d/users/%d", projectID, userID)); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"status": "success"})
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().Int64Var(&userID, "user-id", 0, "User ID")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("user-id")

	return cmd
}
