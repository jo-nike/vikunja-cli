package settings

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newAvatarCmd() *cobra.Command {
	var provider string

	cmd := &cobra.Command{
		Use:   "avatar",
		Short: "Set avatar provider or upload avatar",
		Long:  "Set the avatar provider (gravatar, initials, upload, marble) or upload an avatar file with --file.",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{"avatar_provider": provider}
			if err := c.Post("/user/settings/avatar", body, nil); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "avatar provider updated"})
		},
	}
	cmd.Flags().StringVar(&provider, "provider", "", "Avatar provider: gravatar, initials, upload, marble (required)")
	cmd.MarkFlagRequired("provider")
	return cmd
}
