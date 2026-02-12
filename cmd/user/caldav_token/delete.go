package caldav_token

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	var id int64

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a CalDAV token",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			if err := c.Delete(fmt.Sprintf("/user/settings/token/caldav/%d", id)); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "CalDAV token deleted"})
		},
	}
	cmd.Flags().Int64Var(&id, "id", 0, "Token ID (required)")
	cmd.MarkFlagRequired("id")
	return cmd
}
