package cmd

import (
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		output.Result(map[string]string{
			"version": Version,
			"commit":  Commit,
			"date":    Date,
		})
	},
}
