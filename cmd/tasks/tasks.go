package tasks

import (
	"github.com/jo-nike/vikunja-cli/cmd/tasks/assignees"
	"github.com/jo-nike/vikunja-cli/cmd/tasks/attachments"
	"github.com/jo-nike/vikunja-cli/cmd/tasks/comments"
	tasklabels "github.com/jo-nike/vikunja-cli/cmd/tasks/labels"
	"github.com/jo-nike/vikunja-cli/cmd/tasks/relations"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tasks",
		Short: "Manage tasks",
	}
	cmd.AddCommand(newListCmd())
	cmd.AddCommand(newCreateCmd())
	cmd.AddCommand(newGetCmd())
	cmd.AddCommand(newUpdateCmd())
	cmd.AddCommand(newDeleteCmd())
	cmd.AddCommand(newBulkCmd())
	cmd.AddCommand(newMoveCmd())
	cmd.AddCommand(assignees.NewCmd())
	cmd.AddCommand(tasklabels.NewCmd())
	cmd.AddCommand(comments.NewCmd())
	cmd.AddCommand(attachments.NewCmd())
	cmd.AddCommand(relations.NewCmd())
	return cmd
}
