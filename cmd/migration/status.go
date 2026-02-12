package migration

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newStatusCmd() *cobra.Command {
	var service string

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get migration status for a service",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var result interface{}
			if err := c.Get(fmt.Sprintf("/migration/%s/status", service), &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().StringVar(&service, "service", "", "Migration service (required)")
	cmd.MarkFlagRequired("service")
	return cmd
}
