package users

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newUpdateCmd() *cobra.Command {
	var projectID, userID int64
	var right int

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a user's right on a project",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			body := map[string]interface{}{
				"right": right,
			}

			var result map[string]interface{}
			if err := c.Update(fmt.Sprintf("/projects/%d/users/%d", projectID, userID), body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().Int64Var(&userID, "user-id", 0, "User ID")
	cmd.Flags().IntVar(&right, "right", 0, "Right (0=read, 1=read&write, 2=admin)")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("user-id")
	cmd.MarkFlagRequired("right")

	return cmd
}
