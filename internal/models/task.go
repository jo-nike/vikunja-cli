package models

type Task struct {
	ID              int64    `json:"id"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	Done            bool     `json:"done"`
	DoneAt          string   `json:"done_at"`
	DueDate         string   `json:"due_date"`
	StartDate       string   `json:"start_date"`
	EndDate         string   `json:"end_date"`
	Priority        int      `json:"priority"`
	PercentDone     float64  `json:"percent_done"`
	HexColor        string   `json:"hex_color"`
	ProjectID       int64    `json:"project_id"`
	RepeatAfter     int64    `json:"repeat_after"`
	RepeatMode      int      `json:"repeat_mode"`
	Position        float64  `json:"position"`
	BucketID        int64    `json:"bucket_id"`
	CoverImageAttachmentID int64 `json:"cover_image_attachment_id"`
	IsFavorite      bool     `json:"is_favorite"`
	Assignees       []User   `json:"assignees"`
	Labels          []Label  `json:"labels"`
	Created         string   `json:"created"`
	Updated         string   `json:"updated"`
	CreatedBy       *User    `json:"created_by"`
}

type BulkTask struct {
	TaskIDs []int64 `json:"task_ids"`
}

type TaskRelation struct {
	TaskID           int64  `json:"task_id"`
	OtherTaskID      int64  `json:"other_task_id"`
	RelationKind     string `json:"relation_kind"`
	Created          string `json:"created"`
	CreatedBy        *User  `json:"created_by"`
}
