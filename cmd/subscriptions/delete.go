package subscriptions

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	var entityType string
	var entityID int64

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Unsubscribe from an entity",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			path := fmt.Sprintf("/subscriptions/%s/%d", entityType, entityID)
			if err := c.Delete(path); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "unsubscribed"})
		},
	}
	cmd.Flags().StringVar(&entityType, "entity-type", "", "Entity type: project or task (required)")
	cmd.Flags().Int64Var(&entityID, "entity-id", 0, "Entity ID (required)")
	cmd.MarkFlagRequired("entity-type")
	cmd.MarkFlagRequired("entity-id")
	return cmd
}
