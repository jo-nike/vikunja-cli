package models

type Notification struct {
	ID         int64       `json:"id"`
	Name       string      `json:"name"`
	Notification interface{} `json:"notification"`
	Read       bool        `json:"read"`
	ReadAt     string      `json:"read_at"`
	Created    string      `json:"created"`
}
