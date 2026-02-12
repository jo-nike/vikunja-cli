package export

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newDownloadCmd() *cobra.Command {
	var outputPath string
	var password string

	cmd := &cobra.Command{
		Use:   "download",
		Short: "Download the data export",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			body := map[string]interface{}{}
			if password != "" {
				body["password"] = password
			}
			if err := c.PostDownloadFile("/user/export/download", outputPath, body); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "export downloaded to " + outputPath})
		},
	}
	cmd.Flags().StringVar(&outputPath, "output", "vikunja-export.zip", "Output file path")
	cmd.Flags().StringVar(&password, "password", "", "Account password for confirmation")
	return cmd
}
