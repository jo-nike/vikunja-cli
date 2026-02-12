package system

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "system",
		Short: "System information",
	}
	cmd.AddCommand(newInfoCmd())
	return cmd
}

func newInfoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Get system info from the Vikunja instance",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			var result map[string]interface{}
			if err := c.Get("/info", &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
}
