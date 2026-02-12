package reactions

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	var taskID, commentID int64
	var value string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Remove a reaction from a task or comment",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var path string
			if commentID > 0 {
				path = fmt.Sprintf("/tasks/%d/comments/%d/reactions/delete", taskID, commentID)
			} else {
				path = fmt.Sprintf("/tasks/%d/reactions/delete", taskID)
			}
			body := map[string]interface{}{"value": value}
			if err := c.Post(path, body, nil); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "reaction removed"})
		},
	}
	cmd.Flags().Int64Var(&taskID, "task-id", 0, "Task ID (required)")
	cmd.Flags().Int64Var(&commentID, "comment-id", 0, "Comment ID (for comment reactions)")
	cmd.Flags().StringVar(&value, "value", "", "Reaction value/emoji (required)")
	cmd.MarkFlagRequired("task-id")
	cmd.MarkFlagRequired("value")
	return cmd
}
