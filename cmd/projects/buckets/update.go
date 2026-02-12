package buckets

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newUpdateCmd() *cobra.Command {
	var projectID, viewID, id int64
	var title string
	var limit int
	var position float64

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a bucket in a project view",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			body := map[string]interface{}{}
			if cmd.Flags().Changed("title") {
				body["title"] = title
			}
			if cmd.Flags().Changed("limit") {
				body["limit"] = limit
			}
			if cmd.Flags().Changed("position") {
				body["position"] = position
			}

			var bucket models.Bucket
			if err := c.Update(fmt.Sprintf("/projects/%d/views/%d/buckets/%d", projectID, viewID, id), body, &bucket); err != nil {
				output.Error(err)
			}
			output.Result(bucket)
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().Int64Var(&viewID, "view-id", 0, "View ID")
	cmd.Flags().Int64Var(&id, "id", 0, "Bucket ID")
	cmd.Flags().StringVar(&title, "title", "", "Bucket title")
	cmd.Flags().IntVar(&limit, "limit", 0, "Maximum number of tasks in this bucket")
	cmd.Flags().Float64Var(&position, "position", 0, "Bucket position")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("view-id")
	cmd.MarkFlagRequired("id")

	return cmd
}
