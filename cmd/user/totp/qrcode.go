package totp

import (
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newQRCodeCmd() *cobra.Command {
	var outputPath string

	cmd := &cobra.Command{
		Use:   "qrcode",
		Short: "Download TOTP QR code image",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()
			if err := c.DownloadFile("/user/settings/totp/qrcode", outputPath); err != nil {
				output.Error(err)
			}
			output.Result(map[string]string{"message": "QR code saved to " + outputPath})
		},
	}
	cmd.Flags().StringVar(&outputPath, "output", "totp-qrcode.png", "Output file path")
	return cmd
}
