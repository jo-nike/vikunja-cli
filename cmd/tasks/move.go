package tasks

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/jo-nike/vikunja-cli/internal/resolve"
	"github.com/spf13/cobra"
)

func newMoveCmd() *cobra.Command {
	var id int64
	var bucketID int64
	var bucketName string
	var projectID, viewID int64

	cmd := &cobra.Command{
		Use:   "move",
		Short: "Move a task to a different bucket",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			// Resolve project ID from the task if not provided
			if projectID == 0 {
				pid, err := resolve.TaskProjectID(c, id)
				if err != nil {
					output.Error(err)
				}
				projectID = pid
			}

			// Resolve kanban view if not provided
			if viewID == 0 {
				view, err := resolve.FindKanbanView(c, projectID)
				if err != nil {
					output.Error(err)
				}
				viewID = view.ID
			}

			// Resolve bucket name to ID
			if cmd.Flags().Changed("bucket") {
				bucket, err := resolve.BucketByName(c, projectID, viewID, bucketName)
				if err != nil {
					output.Error(err)
				}
				bucketID = bucket.ID
			}

			if bucketID == 0 {
				output.Error(fmt.Errorf("--bucket or --bucket-id is required"))
			}

			// Use the correct API endpoint for moving tasks between buckets
			path := fmt.Sprintf("/projects/%d/views/%d/buckets/%d/tasks", projectID, viewID, bucketID)
			body := models.TaskBucketAssignment{TaskID: id}
			var result models.TaskBucketAssignment
			if err := c.Update(path, body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&id, "id", 0, "Task ID (required)")
	cmd.Flags().Int64Var(&bucketID, "bucket-id", 0, "Target bucket ID")
	cmd.Flags().StringVar(&bucketName, "bucket", "", "Target bucket name (resolved to ID)")
	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID (skips extra API call to infer project)")
	cmd.Flags().Int64Var(&viewID, "view-id", 0, "View ID (for bucket name disambiguation)")
	_ = cmd.MarkFlagRequired("id")
	cmd.MarkFlagsMutuallyExclusive("bucket-id", "bucket")
	return cmd
}
