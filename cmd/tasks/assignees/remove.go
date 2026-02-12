package assignees

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newRemoveCmd() *cobra.Command {
	var taskID, userID int64

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove an assignee from a task",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			path := fmt.Sprintf("/tasks/%d/assignees/%d", taskID, userID)
			if err := c.Delete(path); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "assignee removed successfully"})
		},
	}
	cmd.Flags().Int64Var(&taskID, "task-id", 0, "Task ID (required)")
	cmd.Flags().Int64Var(&userID, "user-id", 0, "User ID to remove (required)")
	_ = cmd.MarkFlagRequired("task-id")
	_ = cmd.MarkFlagRequired("user-id")
	return cmd
}
