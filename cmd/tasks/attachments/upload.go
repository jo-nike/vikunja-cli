package attachments

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newUploadCmd() *cobra.Command {
	var taskID int64
	var filePath string

	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload an attachment to a task",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			path := fmt.Sprintf("/tasks/%d/attachments", taskID)
			var result models.TaskAttachment
			if err := c.DoUpload("PUT", path, "files", filePath, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&taskID, "task-id", 0, "Task ID (required)")
	cmd.Flags().StringVar(&filePath, "file", "", "Path to the file to upload (required)")
	_ = cmd.MarkFlagRequired("task-id")
	_ = cmd.MarkFlagRequired("file")
	return cmd
}
