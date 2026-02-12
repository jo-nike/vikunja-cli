package shares

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	var projectID int64
	var right, sharingType int

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a share link for a project",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			body := map[string]interface{}{
				"right":        right,
				"sharing_type": sharingType,
			}

			var share models.Share
			if err := c.Create(fmt.Sprintf("/projects/%d/shares", projectID), body, &share); err != nil {
				output.Error(err)
			}
			output.Result(share)
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().IntVar(&right, "right", 0, "Share right (0=read, 1=read&write, 2=admin)")
	cmd.Flags().IntVar(&sharingType, "sharing-type", 0, "Sharing type (0=without password, 1=with password)")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("right")
	cmd.MarkFlagRequired("sharing-type")

	return cmd
}
