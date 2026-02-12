package assignees

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newAddCmd() *cobra.Command {
	var taskID, userID int64

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add an assignee to a task",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			path := fmt.Sprintf("/tasks/%d/assignees", taskID)
			body := map[string]interface{}{
				"user_id": userID,
			}
			var result models.Task
			if err := c.Create(path, body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&taskID, "task-id", 0, "Task ID (required)")
	cmd.Flags().Int64Var(&userID, "user-id", 0, "User ID to assign (required)")
	_ = cmd.MarkFlagRequired("task-id")
	_ = cmd.MarkFlagRequired("user-id")
	return cmd
}
