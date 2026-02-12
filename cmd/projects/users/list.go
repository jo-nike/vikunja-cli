package users

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	var projectID int64
	var page, perPage int

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List user members of a project",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			opts := cmdutil.PaginationOpts(page, perPage)

			var users []models.User
			info, err := c.GetList(fmt.Sprintf("/projects/%d/users", projectID), &users, opts...)
			if err != nil {
				output.Error(err)
			}
			output.ResultList(users, info)
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().IntVar(&page, "page", 0, "Page number")
	cmd.Flags().IntVar(&perPage, "per-page", 0, "Number of items per page")
	cmd.MarkFlagRequired("project-id")

	return cmd
}
