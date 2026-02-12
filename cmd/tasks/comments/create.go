package comments

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	var taskID int64
	var comment string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a comment on a task",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			path := fmt.Sprintf("/tasks/%d/comments", taskID)
			body := map[string]interface{}{
				"comment": comment,
			}
			var result models.TaskComment
			if err := c.Create(path, body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&taskID, "task-id", 0, "Task ID (required)")
	cmd.Flags().StringVar(&comment, "comment", "", "Comment text (required)")
	_ = cmd.MarkFlagRequired("task-id")
	_ = cmd.MarkFlagRequired("comment")
	return cmd
}
