package notifications

import (
	"github.com/jo-nike/vikunja-cli/internal/client"
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	var page, perPage int

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List notifications",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var opts []client.RequestOption
			if page > 0 {
				opts = append(opts, client.WithPage(page))
			}
			if perPage > 0 {
				opts = append(opts, client.WithPerPage(perPage))
			}
			var result []models.Notification
			info, err := c.GetList("/notifications", &result, opts...)
			if err != nil {
				output.Error(err)
			}
			output.ResultList(result, info)
		},
	}
	cmd.Flags().IntVar(&page, "page", 0, "Page number")
	cmd.Flags().IntVar(&perPage, "per-page", 0, "Items per page")
	return cmd
}
