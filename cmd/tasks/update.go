package tasks

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/jo-nike/vikunja-cli/internal/resolve"
	"github.com/spf13/cobra"
)

func newUpdateCmd() *cobra.Command {
	var id int64
	var title, description string
	var done bool
	var dueDate, startDate, endDate string
	var priority int
	var hexColor string
	var percentDone float64
	var repeatAfter int64
	var bucketID int64
	var bucketName string
	var projectID, viewID int64
	var isFavorite bool

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a task",
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

			body := map[string]interface{}{}

			if cmd.Flags().Changed("title") {
				body["title"] = title
			}
			if cmd.Flags().Changed("description") {
				body["description"] = description
			}
			if cmd.Flags().Changed("done") {
				body["done"] = done
			}
			if cmd.Flags().Changed("due-date") {
				body["due_date"] = dueDate
			}
			if cmd.Flags().Changed("start-date") {
				body["start_date"] = startDate
			}
			if cmd.Flags().Changed("end-date") {
				body["end_date"] = endDate
			}
			if cmd.Flags().Changed("priority") {
				body["priority"] = priority
			}
			if cmd.Flags().Changed("hex-color") {
				body["hex_color"] = hexColor
			}
			if cmd.Flags().Changed("percent-done") {
				body["percent_done"] = percentDone
			}
			if cmd.Flags().Changed("repeat-after") {
				body["repeat_after"] = repeatAfter
			}
			if cmd.Flags().Changed("bucket-id") || cmd.Flags().Changed("bucket") {
				body["bucket_id"] = bucketID
			}
			if cmd.Flags().Changed("is-favorite") {
				body["is_favorite"] = isFavorite
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
	cmd.Flags().StringVar(&title, "title", "", "Task title")
	cmd.Flags().StringVar(&description, "description", "", "Task description")
	cmd.Flags().BoolVar(&done, "done", false, "Mark task as done")
	cmd.Flags().StringVar(&dueDate, "due-date", "", "Due date (RFC3339)")
	cmd.Flags().StringVar(&startDate, "start-date", "", "Start date (RFC3339)")
	cmd.Flags().StringVar(&endDate, "end-date", "", "End date (RFC3339)")
	cmd.Flags().IntVar(&priority, "priority", 0, "Priority (0-5)")
	cmd.Flags().StringVar(&hexColor, "hex-color", "", "Hex color code")
	cmd.Flags().Float64Var(&percentDone, "percent-done", 0, "Percent done (0.0-1.0)")
	cmd.Flags().Int64Var(&repeatAfter, "repeat-after", 0, "Repeat after N seconds")
	cmd.Flags().Int64Var(&bucketID, "bucket-id", 0, "Bucket ID")
	cmd.Flags().StringVar(&bucketName, "bucket", "", "Bucket name (resolved to ID; mutually exclusive with --bucket-id)")
	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID (skips fetching task to infer project)")
	cmd.Flags().Int64Var(&viewID, "view-id", 0, "View ID (for bucket name disambiguation)")
	cmd.Flags().BoolVar(&isFavorite, "is-favorite", false, "Mark as favorite")
	_ = cmd.MarkFlagRequired("id")
	cmd.MarkFlagsMutuallyExclusive("bucket-id", "bucket")
	return cmd
}
