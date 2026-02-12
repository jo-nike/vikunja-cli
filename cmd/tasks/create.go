package tasks

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/jo-nike/vikunja-cli/internal/resolve"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	var projectID int64
	var title, description string
	var done bool
	var dueDate, startDate, endDate string
	var priority int
	var hexColor string
	var percentDone float64
	var repeatAfter int64
	var bucketID int64
	var bucketName string
	var viewID int64
	var isFavorite bool

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new task in a project",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			if cmd.Flags().Changed("bucket") {
				bucket, _, err := resolve.BucketByNameAutoView(c, projectID, viewID, bucketName)
				if err != nil {
					output.Error(err)
				}
				bucketID = bucket.ID
			}

			body := map[string]interface{}{
				"title": title,
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

			path := fmt.Sprintf("/projects/%d/tasks", projectID)
			var result models.Task
			if err := c.Create(path, body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID (required)")
	cmd.Flags().StringVar(&title, "title", "", "Task title (required)")
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
	cmd.Flags().Int64Var(&viewID, "view-id", 0, "View ID (for bucket name disambiguation)")
	cmd.Flags().BoolVar(&isFavorite, "is-favorite", false, "Mark as favorite")
	_ = cmd.MarkFlagRequired("project-id")
	_ = cmd.MarkFlagRequired("title")
	cmd.MarkFlagsMutuallyExclusive("bucket-id", "bucket")
	return cmd
}
