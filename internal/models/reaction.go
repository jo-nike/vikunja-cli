package models

type Reaction struct {
	Value     string `json:"value"`
	CreatedBy *User  `json:"created_by"`
	Created   string `json:"created"`
}
