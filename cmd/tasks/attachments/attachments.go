package attachments

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attachments",
		Short: "Manage task attachments",
	}
	cmd.AddCommand(newListCmd())
	cmd.AddCommand(newUploadCmd())
	cmd.AddCommand(newGetCmd())
	cmd.AddCommand(newDeleteCmd())
	return cmd
}
