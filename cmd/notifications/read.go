package notifications

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newReadCmd() *cobra.Command {
	var id int64

	cmd := &cobra.Command{
		Use:   "read",
		Short: "Mark a notification as read",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			if err := c.Post(fmt.Sprintf("/notifications/%d", id), nil, nil); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "notification marked as read"})
		},
	}
	cmd.Flags().Int64Var(&id, "id", 0, "Notification ID (required)")
	cmd.MarkFlagRequired("id")
	return cmd
}
