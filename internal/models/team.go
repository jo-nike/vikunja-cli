package models

type Team struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Members     []TeamMember `json:"members"`
	CreatedBy   *User        `json:"created_by"`
	Created     string       `json:"created"`
	Updated     string       `json:"updated"`
}

type TeamMember struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Admin   bool   `json:"admin"`
	Created string `json:"created"`
}
