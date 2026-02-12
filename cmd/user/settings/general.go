package settings

import (
	"encoding/json"
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newGeneralCmd() *cobra.Command {
	var name, language, timezone, frontendSettingsJSON string
	var emailReminders, discoverByName, discoverByEmail, overdueReminders bool
	var defaultProjectID int64
	var weekStart int

	cmd := &cobra.Command{
		Use:   "general",
		Short: "Get or update general settings",
		Long:  "Without flags, returns current settings. With flags, updates them.",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			hasChanges := false
			body := map[string]interface{}{}
			for _, f := range []string{"name", "language", "timezone", "email-reminders", "discover-by-name", "discover-by-email", "overdue-reminders", "default-project-id", "week-start", "frontend-settings"} {
				if cmd.Flags().Changed(f) {
					hasChanges = true
					break
				}
			}

			if !hasChanges {
				// API only supports POST; send empty body to retrieve current settings
				var result map[string]interface{}
				if err := c.Post("/user/settings/general", map[string]interface{}{}, &result); err != nil {
					output.Error(err)
				}
				output.Result(result)
				return
			}

			if cmd.Flags().Changed("name") {
				body["name"] = name
			}
			if cmd.Flags().Changed("language") {
				body["language"] = language
			}
			if cmd.Flags().Changed("timezone") {
				body["timezone"] = timezone
			}
			if cmd.Flags().Changed("email-reminders") {
				body["email_reminders_enabled"] = emailReminders
			}
			if cmd.Flags().Changed("discover-by-name") {
				body["discoverable_by_name"] = discoverByName
			}
			if cmd.Flags().Changed("discover-by-email") {
				body["discoverable_by_email"] = discoverByEmail
			}
			if cmd.Flags().Changed("overdue-reminders") {
				body["overdue_tasks_reminders_enabled"] = overdueReminders
			}
			if cmd.Flags().Changed("default-project-id") {
				body["default_project_id"] = defaultProjectID
			}
			if cmd.Flags().Changed("week-start") {
				body["week_start"] = weekStart
			}
			if cmd.Flags().Changed("frontend-settings") {
				var fs interface{}
				if err := json.Unmarshal([]byte(frontendSettingsJSON), &fs); err != nil {
					output.Error(fmt.Errorf("invalid frontend-settings JSON: %w", err))
				}
				body["frontend_settings"] = fs
			}

			var result map[string]interface{}
			if err := c.Post("/user/settings/general", body, &result); err != nil {
				output.Error(err)
			}
			output.Result(result)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Display name")
	cmd.Flags().StringVar(&language, "language", "", "Language code")
	cmd.Flags().StringVar(&timezone, "timezone", "", "Timezone")
	cmd.Flags().BoolVar(&emailReminders, "email-reminders", false, "Enable email reminders")
	cmd.Flags().BoolVar(&discoverByName, "discover-by-name", false, "Discoverable by name")
	cmd.Flags().BoolVar(&discoverByEmail, "discover-by-email", false, "Discoverable by email")
	cmd.Flags().BoolVar(&overdueReminders, "overdue-reminders", false, "Enable overdue task reminders")
	cmd.Flags().Int64Var(&defaultProjectID, "default-project-id", 0, "Default project ID")
	cmd.Flags().IntVar(&weekStart, "week-start", 0, "Week start day (0=Sunday, 1=Monday)")
	cmd.Flags().StringVar(&frontendSettingsJSON, "frontend-settings", "", "Frontend settings as JSON")
	return cmd
}
