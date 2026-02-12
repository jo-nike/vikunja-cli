package buckets

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	var projectID, viewID, id int64

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a bucket from a project view",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			if err := c.Delete(fmt.Sprintf("/projects/%d/views/%d/buckets/%d", projectID, viewID, id)); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"status": "success"})
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().Int64Var(&viewID, "view-id", 0, "View ID")
	cmd.Flags().Int64Var(&id, "id", 0, "Bucket ID")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("view-id")
	cmd.MarkFlagRequired("id")

	return cmd
}
