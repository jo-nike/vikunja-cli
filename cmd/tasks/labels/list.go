package tasklabels

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/client"
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	var taskID int64
	var page, perPage int

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List labels for a task",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var opts []client.RequestOption
			if page > 0 {
				opts = append(opts, client.WithPage(page))
			}
			if perPage > 0 {
				opts = append(opts, client.WithPerPage(perPage))
			}

			path := fmt.Sprintf("/tasks/%d/labels", taskID)
			var result []models.Label
			info, err := c.GetList(path, &result, opts...)
			if err != nil {
				output.Error(err)
			}
			output.ResultList(result, info)
		},
	}
	cmd.Flags().Int64Var(&taskID, "task-id", 0, "Task ID (required)")
	cmd.Flags().IntVar(&page, "page", 0, "Page number")
	cmd.Flags().IntVar(&perPage, "per-page", 0, "Items per page")
	_ = cmd.MarkFlagRequired("task-id")
	return cmd
}
