package models

type Label struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HexColor    string `json:"hex_color"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	CreatedBy   *User  `json:"created_by"`
}

type LabelTask struct {
	LabelID int64 `json:"label_id"`
}

type BulkLabels struct {
	Labels []Label `json:"labels"`
}
