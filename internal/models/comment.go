package models

type TaskComment struct {
	ID      int64  `json:"id"`
	Comment string `json:"comment"`
	Author  *User  `json:"author"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}
