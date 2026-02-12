package views

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	var projectID, id int64

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a project view by ID",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			var view models.ProjectView
			if err := c.Get(fmt.Sprintf("/projects/%d/views/%d", projectID, id), &view); err != nil {
				output.Error(err)
			}
			output.Result(view)
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().Int64Var(&id, "id", 0, "View ID")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("id")

	return cmd
}
