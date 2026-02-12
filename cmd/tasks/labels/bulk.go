package tasklabels

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
	var labelIDs string

	cmd := &cobra.Command{
		Use:   "bulk",
		Short: "Bulk assign labels to a task",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			parts := strings.Split(labelIDs, ",")
			var labels []map[string]int64
			for _, p := range parts {
				p = strings.TrimSpace(p)
				id, err := strconv.ParseInt(p, 10, 64)
				if err != nil {
					output.Error(fmt.Errorf("invalid label ID %q: %w", p, err))
				}
				labels = append(labels, map[string]int64{"id": id})
			}

			path := fmt.Sprintf("/tasks/%d/labels/bulk", taskID)
			body := map[string]interface{}{
				"labels": labels,
			}
			var result models.BulkLabels
			if err := c.Post(path, body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&taskID, "task-id", 0, "Task ID (required)")
	cmd.Flags().StringVar(&labelIDs, "label-ids", "", "Comma-separated label IDs (required)")
	_ = cmd.MarkFlagRequired("task-id")
	_ = cmd.MarkFlagRequired("label-ids")
	return cmd
}
