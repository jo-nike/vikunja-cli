package filters

import (
	"encoding/json"
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newUpdateCmd() *cobra.Command {
	var id int64
	var title, description, filtersJSON string
	var isFavorite bool

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a saved filter",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			// Fetch existing filter to merge with updates
			var existing models.SavedFilter
			if err := c.Get(fmt.Sprintf("/filters/%d", id), &existing); err != nil {
				output.Error(err)
			}

			body := map[string]interface{}{
				"filters": existing.Filters,
			}
			if cmd.Flags().Changed("title") {
				body["title"] = title
			} else {
				body["title"] = existing.Title
			}
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
			if err := c.Update(fmt.Sprintf("/filters/%d", id), body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&id, "id", 0, "Filter ID (required)")
	cmd.MarkFlagRequired("id")
	cmd.Flags().StringVar(&title, "title", "", "Filter title")
	cmd.Flags().StringVar(&description, "description", "", "Filter description")
	cmd.Flags().StringVar(&filtersJSON, "filters", "", "Filter definition as JSON")
	cmd.Flags().BoolVar(&isFavorite, "is-favorite", false, "Mark as favorite")
	return cmd
}
