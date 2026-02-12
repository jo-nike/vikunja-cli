package export

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Manage data export",
	}
	cmd.AddCommand(newRequestCmd())
	cmd.AddCommand(newDownloadCmd())
	cmd.AddCommand(newStatusCmd())
	return cmd
}
