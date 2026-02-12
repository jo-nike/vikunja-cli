package labels

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	var title, description, hexColor string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a label",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{
				"title": title,
			}
			if cmd.Flags().Changed("description") {
				body["description"] = description
			}
			if cmd.Flags().Changed("hex-color") {
				body["hex_color"] = hexColor
			}

			var result models.Label
			if err := c.Create("/labels", body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().StringVar(&title, "title", "", "Label title (required)")
	cmd.MarkFlagRequired("title")
	cmd.Flags().StringVar(&description, "description", "", "Label description")
	cmd.Flags().StringVar(&hexColor, "hex-color", "", "Hex color (e.g. #ff0000)")
	return cmd
}
