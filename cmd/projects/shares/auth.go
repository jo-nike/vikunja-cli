package shares

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newAuthCmd() *cobra.Command {
	var shareHash, password string

	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate against a share",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			body := map[string]interface{}{}
			if cmd.Flags().Changed("password") {
				body["password"] = password
			}

			var auth models.ShareAuth
			if err := c.Post(fmt.Sprintf("/shares/%s/auth", shareHash), body, &auth); err != nil {
				output.Error(err)
			}
			output.Result(auth)
		},
	}

	cmd.Flags().StringVar(&shareHash, "share-hash", "", "Share hash")
	cmd.Flags().StringVar(&password, "password", "", "Share password")
	cmd.MarkFlagRequired("share-hash")

	return cmd
}
