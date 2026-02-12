package models

type Share struct {
	ID        int64  `json:"id"`
	Hash      string `json:"hash"`
	Right     int    `json:"right"`
	SharedBy  *User  `json:"shared_by"`
	SharingType int  `json:"sharing_type"`
	ProjectID int64  `json:"project_id"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
}

type ShareAuth struct {
	Token string `json:"token"`
}
