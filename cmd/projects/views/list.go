package views

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	var projectID int64

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List views for a project",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			var views []models.ProjectView
			if err := c.Get(fmt.Sprintf("/projects/%d/views", projectID), &views); err != nil {
				output.Error(err)
			}
			output.Result(views)
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.MarkFlagRequired("project-id")

	return cmd
}
