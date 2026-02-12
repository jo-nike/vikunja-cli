package buckets

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	var projectID, viewID int64
	var title string
	var limit int

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a bucket in a project view",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			body := map[string]interface{}{
				"title": title,
			}
			if cmd.Flags().Changed("limit") {
				body["limit"] = limit
			}

			var bucket models.Bucket
			if err := c.Create(fmt.Sprintf("/projects/%d/views/%d/buckets", projectID, viewID), body, &bucket); err != nil {
				output.Error(err)
			}
			output.Result(bucket)
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().Int64Var(&viewID, "view-id", 0, "View ID")
	cmd.Flags().StringVar(&title, "title", "", "Bucket title")
	cmd.Flags().IntVar(&limit, "limit", 0, "Maximum number of tasks in this bucket")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("view-id")
	cmd.MarkFlagRequired("title")

	return cmd
}
