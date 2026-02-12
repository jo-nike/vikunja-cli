package projects

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	var title, description, hexColor, identifier string
	var parentProjectID int64

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new project",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			body := map[string]interface{}{}
			if cmd.Flags().Changed("title") {
				body["title"] = title
			}
			if cmd.Flags().Changed("description") {
				body["description"] = description
			}
			if cmd.Flags().Changed("parent-project-id") {
				body["parent_project_id"] = parentProjectID
			}
			if cmd.Flags().Changed("hex-color") {
				body["hex_color"] = hexColor
			}
			if cmd.Flags().Changed("identifier") {
				body["identifier"] = identifier
			}

			var project models.Project
			if err := c.Create("/projects", body, &project); err != nil {
				output.Error(err)
			}
			output.Result(project)
		},
	}

	cmd.Flags().StringVar(&title, "title", "", "Project title")
	cmd.Flags().StringVar(&description, "description", "", "Project description")
	cmd.Flags().Int64Var(&parentProjectID, "parent-project-id", 0, "Parent project ID")
	cmd.Flags().StringVar(&hexColor, "hex-color", "", "Hex color (e.g. #ff0000)")
	cmd.Flags().StringVar(&identifier, "identifier", "", "Project identifier")
	cmd.MarkFlagRequired("title")

	return cmd
}
