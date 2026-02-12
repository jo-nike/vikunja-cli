package projects

import (
	"github.com/jo-nike/vikunja-cli/internal/client"
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	var page, perPage int
	var search string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all projects",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			opts := cmdutil.PaginationOpts(page, perPage)
			if search != "" {
				opts = append(opts, client.WithSearch(search))
			}

			var projects []models.Project
			info, err := c.GetList("/projects", &projects, opts...)
			if err != nil {
				output.Error(err)
			}
			output.ResultList(projects, info)
		},
	}

	cmd.Flags().IntVar(&page, "page", 0, "Page number")
	cmd.Flags().IntVar(&perPage, "per-page", 0, "Number of items per page")
	cmd.Flags().StringVar(&search, "search", "", "Search projects by title")

	return cmd
}
