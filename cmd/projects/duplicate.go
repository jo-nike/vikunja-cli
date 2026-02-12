package projects

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newDuplicateCmd() *cobra.Command {
	var id int64

	cmd := &cobra.Command{
		Use:   "duplicate",
		Short: "Duplicate a project",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			var result models.ProjectDuplicate
			if err := c.Create(fmt.Sprintf("/projects/%d/duplicate", id), nil, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}

	cmd.Flags().Int64Var(&id, "id", 0, "Project ID to duplicate")
	cmd.MarkFlagRequired("id")

	return cmd
}
