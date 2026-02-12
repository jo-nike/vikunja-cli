package tasks

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	var id int64

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a task",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			path := fmt.Sprintf("/tasks/%d", id)
			if err := c.Delete(path); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "task deleted successfully"})
		},
	}
	cmd.Flags().Int64Var(&id, "id", 0, "Task ID (required)")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}
