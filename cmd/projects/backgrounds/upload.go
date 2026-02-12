package backgrounds

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newUploadCmd() *cobra.Command {
	var projectID int64
	var filePath string

	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload a project background image",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			var result models.Project
			if err := c.DoUpload("PUT", fmt.Sprintf("/projects/%d/backgrounds/upload", projectID), "background", filePath, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().StringVar(&filePath, "file", "", "Path to the background image file")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("file")

	return cmd
}
