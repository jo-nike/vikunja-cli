package tasklabels

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newAddCmd() *cobra.Command {
	var taskID, labelID int64

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a label to a task",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			path := fmt.Sprintf("/tasks/%d/labels", taskID)
			body := map[string]interface{}{
				"label_id": labelID,
			}
			var result models.LabelTask
			if err := c.Create(path, body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&taskID, "task-id", 0, "Task ID (required)")
	cmd.Flags().Int64Var(&labelID, "label-id", 0, "Label ID to add (required)")
	_ = cmd.MarkFlagRequired("task-id")
	_ = cmd.MarkFlagRequired("label-id")
	return cmd
}
