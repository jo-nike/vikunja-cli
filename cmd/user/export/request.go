package export

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newRequestCmd() *cobra.Command {
	var password string

	cmd := &cobra.Command{
		Use:   "request",
		Short: "Request a data export",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{"password": password}
			if err := c.Post("/user/export/request", body, nil); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "export requested"})
		},
	}
	cmd.Flags().StringVar(&password, "password", "", "Account password (required)")
	cmd.MarkFlagRequired("password")
	return cmd
}
