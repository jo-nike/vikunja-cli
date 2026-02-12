package tokens

import (
	"encoding/json"
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	var title, expiresAt, permissionsJSON string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an API token",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{"title": title}
			if cmd.Flags().Changed("expires-at") {
				body["expires_at"] = expiresAt
			}
			if cmd.Flags().Changed("permissions") {
				var p interface{}
				if err := json.Unmarshal([]byte(permissionsJSON), &p); err != nil {
					output.Error(fmt.Errorf("invalid permissions JSON: %w", err))
				}
				body["permissions"] = p
			}
			var result models.APIToken
			if err := c.Create("/tokens", body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().StringVar(&title, "title", "", "Token title (required)")
	cmd.MarkFlagRequired("title")
	cmd.Flags().StringVar(&expiresAt, "expires-at", "", "Expiration date (RFC3339)")
	cmd.Flags().StringVar(&permissionsJSON, "permissions", "", "Permissions as JSON")
	return cmd
}
