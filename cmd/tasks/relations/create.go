package relations

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	var taskID, otherTaskID int64
	var kind string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a relation between two tasks",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			path := fmt.Sprintf("/tasks/%d/relations", taskID)
			body := map[string]interface{}{
				"other_task_id": otherTaskID,
				"relation_kind": kind,
			}
			var result models.TaskRelation
			if err := c.Create(path, body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&taskID, "task-id", 0, "Task ID (required)")
	cmd.Flags().Int64Var(&otherTaskID, "other-task-id", 0, "Other task ID (required)")
	cmd.Flags().StringVar(&kind, "kind", "", "Relation kind (required)")
	_ = cmd.MarkFlagRequired("task-id")
	_ = cmd.MarkFlagRequired("other-task-id")
	_ = cmd.MarkFlagRequired("kind")
	return cmd
}
