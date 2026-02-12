package tasks

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	var id int64

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a task by ID",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			path := fmt.Sprintf("/tasks/%d", id)
			var result models.Task
			if err := c.Get(path, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&id, "id", 0, "Task ID (required)")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}
