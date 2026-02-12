package projects

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newUpdateCmd() *cobra.Command {
	var id int64
	var title, description, hexColor, identifier string
	var isArchived, isFavorite bool

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a project",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			body := map[string]interface{}{}
			if cmd.Flags().Changed("title") {
				body["title"] = title
			}
			if cmd.Flags().Changed("description") {
				body["description"] = description
			}
			if cmd.Flags().Changed("is-archived") {
				body["is_archived"] = isArchived
			}
			if cmd.Flags().Changed("hex-color") {
				body["hex_color"] = hexColor
			}
			if cmd.Flags().Changed("identifier") {
				body["identifier"] = identifier
			}
			if cmd.Flags().Changed("is-favorite") {
				body["is_favorite"] = isFavorite
			}

			var project models.Project
			if err := c.Update(fmt.Sprintf("/projects/%d", id), body, &project); err != nil {
				output.Error(err)
			}
			output.Result(project)
		},
	}

	cmd.Flags().Int64Var(&id, "id", 0, "Project ID")
	cmd.Flags().StringVar(&title, "title", "", "Project title")
	cmd.Flags().StringVar(&description, "description", "", "Project description")
	cmd.Flags().BoolVar(&isArchived, "is-archived", false, "Archive the project")
	cmd.Flags().StringVar(&hexColor, "hex-color", "", "Hex color (e.g. #ff0000)")
	cmd.Flags().StringVar(&identifier, "identifier", "", "Project identifier")
	cmd.Flags().BoolVar(&isFavorite, "is-favorite", false, "Mark as favorite")
	cmd.MarkFlagRequired("id")

	return cmd
}
