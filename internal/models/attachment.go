package models

type TaskAttachment struct {
	ID        int64  `json:"id"`
	TaskID    int64  `json:"task_id"`
	FileName  string `json:"file_name"`
	FileSize  int64  `json:"file_size"`
	CreatedBy *User  `json:"created_by"`
	Created   string `json:"created"`
}
