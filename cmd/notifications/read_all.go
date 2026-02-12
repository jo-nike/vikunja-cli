package notifications

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newReadAllCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "read-all",
		Short: "Mark all notifications as read",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			if err := c.Post("/notifications", nil, nil); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "all notifications marked as read"})
		},
	}
}
