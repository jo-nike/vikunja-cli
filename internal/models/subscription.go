package models

type Subscription struct {
	ID         int64  `json:"id"`
	EntityType string `json:"entity"`
	EntityID   int64  `json:"entity_id"`
	Created    string `json:"created"`
}
