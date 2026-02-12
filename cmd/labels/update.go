package labels

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newUpdateCmd() *cobra.Command {
	var id int64
	var title, description, hexColor string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a label",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{}
			if cmd.Flags().Changed("title") {
				body["title"] = title
			}
			if cmd.Flags().Changed("description") {
				body["description"] = description
			}
			if cmd.Flags().Changed("hex-color") {
				body["hex_color"] = hexColor
			}

			var result models.Label
			if err := c.Update(fmt.Sprintf("/labels/%d", id), body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().Int64Var(&id, "id", 0, "Label ID (required)")
	cmd.MarkFlagRequired("id")
	cmd.Flags().StringVar(&title, "title", "", "Label title")
	cmd.Flags().StringVar(&description, "description", "", "Label description")
	cmd.Flags().StringVar(&hexColor, "hex-color", "", "Hex color")
	return cmd
}
