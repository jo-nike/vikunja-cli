package tasks

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/jo-nike/vikunja-cli/internal/resolve"
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	var id, projectID, viewID int64

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a task by ID",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			path := fmt.Sprintf("/tasks/%d", id)
			var result models.Task
			if err := c.Get(path, &result); err != nil {
				output.Error(err)
			}

			// Enrich with real bucket_id from kanban view (non-fatal on error)
			realBucketID, err := resolve.TaskBucketID(c, id, projectID, viewID)
			if err == nil && realBucketID > 0 {
				result.BucketID = realBucketID
			}

			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&id, "id", 0, "Task ID (required)")
	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID (skips extra API call to infer project)")
	cmd.Flags().Int64Var(&viewID, "view-id", 0, "Kanban view ID (skips auto-detection)")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}
