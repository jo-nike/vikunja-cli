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

// TaskBucketAssignment represents the assignment of a task to a bucket,
// used by POST /projects/{project}/views/{view}/buckets/{bucket}/tasks
type TaskBucketAssignment struct {
	TaskID        int64 `json:"task_id"`
	BucketID      int64 `json:"bucket_id,omitempty"`
	ProjectViewID int64 `json:"project_view_id,omitempty"`
}

// TaskBucket is a bucket with embedded tasks, as returned by the
// kanban view tasks endpoint: GET /projects/{id}/views/{viewID}/tasks
type TaskBucket struct {
	ID            int64   `json:"id"`
	Title         string  `json:"title"`
	ProjectViewID int64   `json:"project_view_id"`
	Limit         int     `json:"limit"`
	Position      float64 `json:"position"`
	Count         int     `json:"count"`
	Created       string  `json:"created"`
	Updated       string  `json:"updated"`
	CreatedBy     *User   `json:"created_by"`
	Tasks         []Task  `json:"tasks"`
}
