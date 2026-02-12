package webhooks

import (
	"fmt"
	"strings"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	var projectID int64
	var targetURL, events, secret string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a webhook for a project",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			eventList := strings.Split(events, ",")
			for i := range eventList {
				eventList[i] = strings.TrimSpace(eventList[i])
			}

			body := map[string]interface{}{
				"target_url": targetURL,
				"events":     eventList,
			}
			if cmd.Flags().Changed("secret") {
				body["secret"] = secret
			}

			var webhook models.Webhook
			if err := c.Create(fmt.Sprintf("/projects/%d/webhooks", projectID), body, &webhook); err != nil {
				output.Error(err)
			}
			output.Result(webhook)
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().StringVar(&targetURL, "target-url", "", "Webhook target URL")
	cmd.Flags().StringVar(&events, "events", "", "Comma-separated list of events")
	cmd.Flags().StringVar(&secret, "secret", "", "Webhook secret")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("target-url")
	cmd.MarkFlagRequired("events")

	return cmd
}
