package deletion

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newCancelCmd() *cobra.Command {
	var password string

	cmd := &cobra.Command{
		Use:   "cancel",
		Short: "Cancel account deletion",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{"password": password}
			if err := c.Post("/user/deletion/cancel", body, nil); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "deletion cancelled"})
		},
	}
	cmd.Flags().StringVar(&password, "password", "", "Account password (required)")
	cmd.MarkFlagRequired("password")
	return cmd
}
