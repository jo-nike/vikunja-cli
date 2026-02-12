package shares

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	var projectID, id int64

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a project share",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			if err := c.Delete(fmt.Sprintf("/projects/%d/shares/%d", projectID, id)); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"status": "success"})
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().Int64Var(&id, "id", 0, "Share ID")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("id")

	return cmd
}
