package models

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

type UserSettings struct {
	Name                    string            `json:"name"`
	EmailRemindersEnabled   bool              `json:"email_reminders_enabled"`
	DiscoverableByName      bool              `json:"discoverable_by_name"`
	DiscoverableByEmail     bool              `json:"discoverable_by_email"`
	OverdueTasksRemindersEnabled bool         `json:"overdue_tasks_reminders_enabled"`
	OverdueTasksRemindersTime    string       `json:"overdue_tasks_reminders_time"`
	DefaultProjectID        int64             `json:"default_project_id"`
	WeekStart               int               `json:"week_start"`
	Language                string            `json:"language"`
	Timezone                string            `json:"timezone"`
	FrontendSettings        map[string]interface{} `json:"frontend_settings"`
}

type AuthToken struct {
	Token string `json:"token"`
}

type UserDeletion struct {
	ScheduledAt string `json:"scheduled_at"`
}

type TOTPStatus struct {
	Enabled bool   `json:"enabled"`
	URL     string `json:"url"`
}

type CaldavToken struct {
	ID      int64  `json:"id"`
	Token   string `json:"token"`
	Created string `json:"created"`
}
