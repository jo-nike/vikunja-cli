package user

import (
	"github.com/jo-nike/vikunja-cli/internal/client"
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	var search string
	var page, perPage int

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Search/list users",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var opts []client.RequestOption
			if search != "" {
				opts = append(opts, client.WithSearch(search))
			}
			if page > 0 {
				opts = append(opts, client.WithPage(page))
			}
			if perPage > 0 {
				opts = append(opts, client.WithPerPage(perPage))
			}
			var result []models.User
			info, err := c.GetList("/users", &result, opts...)
			if err != nil {
				output.Error(err)
			}
			output.ResultList(result, info)
		},
	}
	cmd.Flags().StringVar(&search, "search", "", "Search query (required)")
	cmd.MarkFlagRequired("search")
	cmd.Flags().IntVar(&page, "page", 0, "Page number")
	cmd.Flags().IntVar(&perPage, "per-page", 0, "Items per page")
	return cmd
}
