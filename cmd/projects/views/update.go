package views

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newUpdateCmd() *cobra.Command {
	var projectID, id int64
	var title, filter string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a project view",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			body := map[string]interface{}{}
			if cmd.Flags().Changed("title") {
				body["title"] = title
			}
			if cmd.Flags().Changed("filter") {
				body["filter"] = filter
			}

			var view models.ProjectView
			if err := c.Update(fmt.Sprintf("/projects/%d/views/%d", projectID, id), body, &view); err != nil {
				output.Error(err)
			}
			output.Result(view)
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().Int64Var(&id, "id", 0, "View ID")
	cmd.Flags().StringVar(&title, "title", "", "View title")
	cmd.Flags().StringVar(&filter, "filter", "", "View filter")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("id")

	return cmd
}
