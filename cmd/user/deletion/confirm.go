package deletion

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newConfirmCmd() *cobra.Command {
	var token string

	cmd := &cobra.Command{
		Use:   "confirm",
		Short: "Confirm account deletion",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{"token": token}
			if err := c.Post("/user/deletion/confirm", body, nil); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "deletion confirmed"})
		},
	}
	cmd.Flags().StringVar(&token, "token", "", "Confirmation token (required)")
	cmd.MarkFlagRequired("token")
	return cmd
}
