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

			if cmd.Flags().Changed("bucket") {
				if projectID == 0 {
					pid, err := resolve.TaskProjectID(c, id)
					if err != nil {
						output.Error(err)
					}
					projectID = pid
				}
				bucket, _, err := resolve.BucketByNameAutoView(c, projectID, viewID, bucketName)
				if err != nil {
					output.Error(err)
				}
				bucketID = bucket.ID
			}

			body := map[string]interface{}{
				"bucket_id": bucketID,
			}

			path := fmt.Sprintf("/tasks/%d", id)
			var result models.Task
			if err := c.Update(path, body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&id, "id", 0, "Task ID (required)")
	cmd.Flags().Int64Var(&bucketID, "bucket-id", 0, "Target bucket ID")
	cmd.Flags().StringVar(&bucketName, "bucket", "", "Target bucket name (resolved to ID)")
	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID (skips fetching task to infer project)")
	cmd.Flags().Int64Var(&viewID, "view-id", 0, "View ID (for bucket name disambiguation)")
	_ = cmd.MarkFlagRequired("id")
	cmd.MarkFlagsMutuallyExclusive("bucket-id", "bucket")
	return cmd
}
