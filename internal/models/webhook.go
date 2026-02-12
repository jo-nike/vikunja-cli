package models

type Webhook struct {
	ID        int64  `json:"id"`
	TargetURL string `json:"target_url"`
	Events    []string `json:"events"`
	ProjectID int64  `json:"project_id"`
	Secret    string `json:"secret"`
	CreatedBy *User  `json:"created_by"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
}
