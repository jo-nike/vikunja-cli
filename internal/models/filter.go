package models

type SavedFilter struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Filters     interface{} `json:"filters"`
	IsFavorite  bool   `json:"is_favorite"`
	Owner       *User  `json:"owner"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
}
