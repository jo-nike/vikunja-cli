package models

type Project struct {
	ID                 int64   `json:"id"`
	Title              string  `json:"title"`
	Description        string  `json:"description"`
	Identifier         string  `json:"identifier"`
	HexColor           string  `json:"hex_color"`
	ParentProjectID    int64   `json:"parent_project_id"`
	Position           float64 `json:"position"`
	IsArchived         bool    `json:"is_archived"`
	IsFavorite         bool    `json:"is_favorite"`
	DefaultBucketID    int64   `json:"default_bucket_id"`
	DoneBucketID       int64   `json:"done_bucket_id"`
	Owner              *User   `json:"owner"`
	Created            string  `json:"created"`
	Updated            string  `json:"updated"`
}

type ProjectDuplicate struct {
	ProjectID       int64    `json:"project_id"`
	DuplicatedProject *Project `json:"duplicated_project"`
}
