package subscriptions

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	var entityType string
	var entityID int64

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Subscribe to an entity",
		Long:  "Subscribe to a project or task. Entity type must be 'project' or 'task'.",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			path := fmt.Sprintf("/subscriptions/%s/%d", entityType, entityID)
			var result models.Subscription
			if err := c.Create(path, nil, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().StringVar(&entityType, "entity-type", "", "Entity type: project or task (required)")
	cmd.Flags().Int64Var(&entityID, "entity-id", 0, "Entity ID (required)")
	cmd.MarkFlagRequired("entity-type")
	cmd.MarkFlagRequired("entity-id")
	return cmd
}
