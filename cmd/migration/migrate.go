package migration

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newMigrateCmd() *cobra.Command {
	var service, code string

	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Start a migration from a service",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{"code": code}
			var result interface{}
			if err := c.Post(fmt.Sprintf("/migration/%s/migrate", service), body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().StringVar(&service, "service", "", "Migration service (required)")
	cmd.Flags().StringVar(&code, "code", "", "Auth code from the migration service (required)")
	cmd.MarkFlagRequired("service")
	cmd.MarkFlagRequired("code")
	return cmd
}
