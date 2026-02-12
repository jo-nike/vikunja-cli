package models

type APIToken struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	Token       string      `json:"token"`
	Permissions interface{} `json:"permissions"`
	ExpiresAt   string      `json:"expires_at"`
	Created     string      `json:"created"`
}

type APIRoute struct {
	Path   string   `json:"path"`
	Method string   `json:"method"`
	Routes []string `json:"routes"`
}
