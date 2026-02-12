package tasklabels

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newRemoveCmd() *cobra.Command {
	var taskID, labelID int64

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a label from a task",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			path := fmt.Sprintf("/tasks/%d/labels/%d", taskID, labelID)
			if err := c.Delete(path); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "label removed successfully"})
		},
	}
	cmd.Flags().Int64Var(&taskID, "task-id", 0, "Task ID (required)")
	cmd.Flags().Int64Var(&labelID, "label-id", 0, "Label ID to remove (required)")
	_ = cmd.MarkFlagRequired("task-id")
	_ = cmd.MarkFlagRequired("label-id")
	return cmd
}
