package backgrounds

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jo-nike/vikunja-cli/internal/client"
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newSearchCmd() *cobra.Command {
	var search string
	var page, perPage int

	cmd := &cobra.Command{
		Use:   "search",
		Short: "Search Unsplash for background images",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			opts := []client.RequestOption{
				client.WithQuery("s", search),
			}
			if page > 0 {
				opts = append(opts, client.WithQuery("p", strconv.Itoa(page)))
			}
			if perPage > 0 {
				opts = append(opts, client.WithPerPage(perPage))
			}

			raw, err := c.GetRaw("/backgrounds/unsplash/search", opts...)
			if err != nil {
				output.Error(err)
			}

			var results []map[string]interface{}
			if err := json.Unmarshal(raw, &results); err != nil {
				output.Error(fmt.Errorf("parsing response: %w", err))
			}
			output.Result(results)
		},
	}

	cmd.Flags().StringVar(&search, "search", "", "Search query")
	cmd.Flags().IntVar(&page, "page", 0, "Page number")
	cmd.Flags().IntVar(&perPage, "per-page", 0, "Number of items per page")
	cmd.MarkFlagRequired("search")

	return cmd
}
