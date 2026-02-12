package views

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	var projectID int64
	var title, viewKind, filter, bucketConfigMode string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a view for a project",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			body := map[string]interface{}{}
			if cmd.Flags().Changed("title") {
				body["title"] = title
			}
			if cmd.Flags().Changed("view-kind") {
				body["view_kind"] = viewKind
			}
			if cmd.Flags().Changed("filter") {
				body["filter"] = filter
			}
			if cmd.Flags().Changed("bucket-config-mode") {
				body["bucket_configuration_mode"] = bucketConfigMode
			}

			var view models.ProjectView
			if err := c.Create(fmt.Sprintf("/projects/%d/views", projectID), body, &view); err != nil {
				output.Error(err)
			}
			output.Result(view)
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().StringVar(&title, "title", "", "View title")
	cmd.Flags().StringVar(&viewKind, "view-kind", "", "View kind (list, gantt, table, kanban)")
	cmd.Flags().StringVar(&filter, "filter", "", "View filter")
	cmd.Flags().StringVar(&bucketConfigMode, "bucket-config-mode", "", "Bucket configuration mode")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("title")
	cmd.MarkFlagRequired("view-kind")

	return cmd
}
