package assignees

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newBulkCmd() *cobra.Command {
	var taskID int64
	var userIDs string

	cmd := &cobra.Command{
		Use:   "bulk",
		Short: "Bulk assign users to a task",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			parts := strings.Split(userIDs, ",")
			var assignees []map[string]int64
			for _, p := range parts {
				p = strings.TrimSpace(p)
				id, err := strconv.ParseInt(p, 10, 64)
				if err != nil {
					output.Error(fmt.Errorf("invalid user ID %q: %w", p, err))
				}
				assignees = append(assignees, map[string]int64{"id": id})
			}

			path := fmt.Sprintf("/tasks/%d/assignees/bulk", taskID)
			body := map[string]interface{}{
				"assignees": assignees,
			}
			var result models.Task
			if err := c.Post(path, body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&taskID, "task-id", 0, "Task ID (required)")
	cmd.Flags().StringVar(&userIDs, "user-ids", "", "Comma-separated user IDs (required)")
	_ = cmd.MarkFlagRequired("task-id")
	_ = cmd.MarkFlagRequired("user-ids")
	return cmd
}
