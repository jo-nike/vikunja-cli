package migration

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newAuthCmd() *cobra.Command {
	var service string

	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Get the auth URL for a migration service",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var result interface{}
			if err := c.Get(fmt.Sprintf("/migration/%s/auth", service), &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().StringVar(&service, "service", "", "Migration service: todoist, trello, microsoft-todo, ticktick (required)")
	cmd.MarkFlagRequired("service")
	return cmd
}
