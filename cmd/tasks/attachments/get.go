package attachments

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	var taskID, id int64
	var outputPath string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Download an attachment",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			path := fmt.Sprintf("/tasks/%d/attachments/%d", taskID, id)
			if err := c.DownloadFile(path, outputPath); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": fmt.Sprintf("downloaded to %s", outputPath)})
		},
	}
	cmd.Flags().Int64Var(&taskID, "task-id", 0, "Task ID (required)")
	cmd.Flags().Int64Var(&id, "id", 0, "Attachment ID (required)")
	cmd.Flags().StringVar(&outputPath, "output", "", "Output file path (required)")
	_ = cmd.MarkFlagRequired("task-id")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("output")
	return cmd
}
