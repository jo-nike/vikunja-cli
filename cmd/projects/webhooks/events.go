package webhooks

import (
	"encoding/json"
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newEventsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "events",
		Short: "List available webhook events",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			raw, err := c.GetRaw("/webhooks/events")
			if err != nil {
				output.Error(err)
			}

			var events []string
			if err := json.Unmarshal(raw, &events); err != nil {
				output.Error(fmt.Errorf("parsing response: %w", err))
			}
			output.Result(events)
		},
	}

	return cmd
}
