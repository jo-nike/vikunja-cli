package backgrounds

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	var projectID int64
	var outputPath string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Download a project background",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			if err := c.DownloadFile(fmt.Sprintf("/projects/%d/background", projectID), outputPath); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"status": "success", "path": outputPath})
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().StringVar(&outputPath, "output", "", "Output file path")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("output")

	return cmd
}
