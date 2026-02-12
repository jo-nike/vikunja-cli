package backgrounds

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newUnsplashCmd() *cobra.Command {
	var projectID int64
	var imageID string

	cmd := &cobra.Command{
		Use:   "unsplash",
		Short: "Set an Unsplash image as project background",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			body := map[string]interface{}{
				"id": imageID,
			}

			var project models.Project
			if err := c.Post(fmt.Sprintf("/projects/%d/backgrounds/unsplash", projectID), body, &project); err != nil {
				output.Error(err)
			}
			output.Result(project)
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().StringVar(&imageID, "image-id", "", "Unsplash image ID")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("image-id")

	return cmd
}
