package projects

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	var id int64

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a project by ID",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			var project models.Project
			if err := c.Get(fmt.Sprintf("/projects/%d", id), &project); err != nil {
				output.Error(err)
			}
			output.Result(project)
		},
	}

	cmd.Flags().Int64Var(&id, "id", 0, "Project ID")
	cmd.MarkFlagRequired("id")

	return cmd
}
