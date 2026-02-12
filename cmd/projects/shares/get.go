package shares

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	var shareHash string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a share by hash",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			var share models.Share
			if err := c.Get(fmt.Sprintf("/shares/%s", shareHash), &share); err != nil {
				output.Error(err)
			}
			output.Result(share)
		},
	}

	cmd.Flags().StringVar(&shareHash, "share-hash", "", "Share hash")
	cmd.MarkFlagRequired("share-hash")

	return cmd
}
