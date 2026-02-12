package filters

import (
	"encoding/json"
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	var title, description, filtersJSON string
	var isFavorite bool

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a saved filter",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{"title": title}
			if cmd.Flags().Changed("description") {
				body["description"] = description
			}
			if cmd.Flags().Changed("filters") {
				var f interface{}
				if err := json.Unmarshal([]byte(filtersJSON), &f); err != nil {
					output.Error(fmt.Errorf("invalid filters JSON: %w", err))
				}
				body["filters"] = f
			}
			if cmd.Flags().Changed("is-favorite") {
				body["is_favorite"] = isFavorite
			}
			var result models.SavedFilter
			if err := c.Create("/filters", body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().StringVar(&title, "title", "", "Filter title (required)")
	cmd.MarkFlagRequired("title")
	cmd.Flags().StringVar(&description, "description", "", "Filter description")
	cmd.Flags().StringVar(&filtersJSON, "filters", "", `Filter definition as JSON (e.g. '{"filter":"done = false","sort_by":["due_date"],"order_by":["asc"]}')`)
	cmd.Flags().BoolVar(&isFavorite, "is-favorite", false, "Mark as favorite")
	return cmd
}
