package projects

import (
	"github.com/jo-nike/vikunja-cli/cmd/projects/backgrounds"
	"github.com/jo-nike/vikunja-cli/cmd/projects/buckets"
	"github.com/jo-nike/vikunja-cli/cmd/projects/shares"
	projectteams "github.com/jo-nike/vikunja-cli/cmd/projects/teams"
	"github.com/jo-nike/vikunja-cli/cmd/projects/users"
	"github.com/jo-nike/vikunja-cli/cmd/projects/views"
	"github.com/jo-nike/vikunja-cli/cmd/projects/webhooks"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projects",
		Short: "Manage projects",
	}
	cmd.AddCommand(newListCmd())
	cmd.AddCommand(newCreateCmd())
	cmd.AddCommand(newGetCmd())
	cmd.AddCommand(newUpdateCmd())
	cmd.AddCommand(newDeleteCmd())
	cmd.AddCommand(newDuplicateCmd())
	cmd.AddCommand(shares.NewCmd())
	cmd.AddCommand(users.NewCmd())
	cmd.AddCommand(projectteams.NewCmd())
	cmd.AddCommand(views.NewCmd())
	cmd.AddCommand(backgrounds.NewCmd())
	cmd.AddCommand(webhooks.NewCmd())
	cmd.AddCommand(buckets.NewCmd())
	return cmd
}
