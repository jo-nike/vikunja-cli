package tasks

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/jo-nike/vikunja-cli/internal/resolve"
	"github.com/spf13/cobra"
)

func newBulkCmd() *cobra.Command {
	var taskIDs string
	var done bool
	var priority int
	var dueDate string
	var bucketID int64
	var bucketName string
	var projectID, viewID int64

	cmd := &cobra.Command{
		Use:   "bulk",
		Short: "Bulk update tasks",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			parts := strings.Split(taskIDs, ",")
			var ids []int64
			for _, p := range parts {
				p = strings.TrimSpace(p)
				id, err := strconv.ParseInt(p, 10, 64)
				if err != nil {
					output.Error(fmt.Errorf("invalid task ID %q: %w", p, err))
				}
				ids = append(ids, id)
			}

			if cmd.Flags().Changed("bucket") {
				if projectID == 0 {
					output.Error(fmt.Errorf("--project-id is required when using --bucket with bulk operations"))
				}
				bucket, _, err := resolve.BucketByNameAutoView(c, projectID, viewID, bucketName)
				if err != nil {
					output.Error(err)
				}
				bucketID = bucket.ID
			}

			body := map[string]interface{}{
				"task_ids": ids,
			}

			if cmd.Flags().Changed("done") {
				body["done"] = done
			}
			if cmd.Flags().Changed("priority") {
				body["priority"] = priority
			}
			if cmd.Flags().Changed("due-date") {
				body["due_date"] = dueDate
			}
			if cmd.Flags().Changed("bucket-id") || cmd.Flags().Changed("bucket") {
				body["bucket_id"] = bucketID
			}

			var result []models.Task
			if err := c.Post("/tasks/bulk", body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().StringVar(&taskIDs, "task-ids", "", "Comma-separated task IDs (required)")
	cmd.Flags().BoolVar(&done, "done", false, "Mark tasks as done")
	cmd.Flags().IntVar(&priority, "priority", 0, "Priority (0-5)")
	cmd.Flags().StringVar(&dueDate, "due-date", "", "Due date (RFC3339)")
	cmd.Flags().Int64Var(&bucketID, "bucket-id", 0, "Target bucket ID")
	cmd.Flags().StringVar(&bucketName, "bucket", "", "Target bucket name (resolved to ID; requires --project-id)")
	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID (required when using --bucket)")
	cmd.Flags().Int64Var(&viewID, "view-id", 0, "View ID (for bucket name disambiguation)")
	_ = cmd.MarkFlagRequired("task-ids")
	cmd.MarkFlagsMutuallyExclusive("bucket-id", "bucket")
	return cmd
}
