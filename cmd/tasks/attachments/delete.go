package attachments

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	var taskID, id int64

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an attachment",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			path := fmt.Sprintf("/tasks/%d/attachments/%d", taskID, id)
			if err := c.Delete(path); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "attachment deleted successfully"})
		},
	}
	cmd.Flags().Int64Var(&taskID, "task-id", 0, "Task ID (required)")
	cmd.Flags().Int64Var(&id, "id", 0, "Attachment ID (required)")
	_ = cmd.MarkFlagRequired("task-id")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}
