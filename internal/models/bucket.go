package models

type Bucket struct {
	ID        int64   `json:"id"`
	Title     string  `json:"title"`
	ProjectViewID int64 `json:"project_view_id"`
	Limit     int     `json:"limit"`
	Position  float64 `json:"position"`
	Count     int     `json:"count"`
	Created   string  `json:"created"`
	Updated   string  `json:"updated"`
	CreatedBy *User   `json:"created_by"`
}
