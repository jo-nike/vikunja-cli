package reactions

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	var taskID, commentID int64

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List reactions on a task or comment",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var path string
			if commentID > 0 {
				path = fmt.Sprintf("/tasks/%d/comments/%d/reactions", taskID, commentID)
			} else {
				path = fmt.Sprintf("/tasks/%d/reactions", taskID)
			}
			var result interface{}
			if err := c.Get(path, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&taskID, "task-id", 0, "Task ID (required)")
	cmd.Flags().Int64Var(&commentID, "comment-id", 0, "Comment ID (for comment reactions)")
	cmd.MarkFlagRequired("task-id")
	return cmd
}
