package webhooks

import (
	"fmt"
	"strings"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newUpdateCmd() *cobra.Command {
	var projectID, id int64
	var targetURL, events, secret string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a project webhook",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			body := map[string]interface{}{}
			if cmd.Flags().Changed("target-url") {
				body["target_url"] = targetURL
			}
			if cmd.Flags().Changed("events") {
				eventList := strings.Split(events, ",")
				for i := range eventList {
					eventList[i] = strings.TrimSpace(eventList[i])
				}
				body["events"] = eventList
			}
			if cmd.Flags().Changed("secret") {
				body["secret"] = secret
			}

			var webhook models.Webhook
			if err := c.Update(fmt.Sprintf("/projects/%d/webhooks/%d", projectID, id), body, &webhook); err != nil {
				output.Error(err)
			}
			output.Result(webhook)
		},
	}

	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().Int64Var(&id, "id", 0, "Webhook ID")
	cmd.Flags().StringVar(&targetURL, "target-url", "", "Webhook target URL")
	cmd.Flags().StringVar(&events, "events", "", "Comma-separated list of events")
	cmd.Flags().StringVar(&secret, "secret", "", "Webhook secret")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("id")

	return cmd
}
