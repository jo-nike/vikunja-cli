package backgrounds

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "backgrounds",
		Short: "Manage project backgrounds",
	}
	cmd.AddCommand(newGetCmd())
	cmd.AddCommand(newDeleteCmd())
	cmd.AddCommand(newUploadCmd())
	cmd.AddCommand(newUnsplashCmd())
	cmd.AddCommand(newSearchCmd())
	return cmd
}
