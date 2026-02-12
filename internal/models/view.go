package models

type ProjectView struct {
	ID               int64       `json:"id"`
	Title            string      `json:"title"`
	ProjectID        int64       `json:"project_id"`
	ViewKind         string      `json:"view_kind"`
	Filter           interface{} `json:"filter"`
	Position         float64     `json:"position"`
	BucketConfigMode string      `json:"bucket_configuration_mode"`
	DefaultBucketID  int64       `json:"default_bucket_id"`
	DoneBucketID     int64       `json:"done_bucket_id"`
	Created          string      `json:"created"`
	Updated          string      `json:"updated"`
}
