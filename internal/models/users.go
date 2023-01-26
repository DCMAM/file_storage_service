package models

// User represents the entity of user
type User struct {
	UserID   int64  `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
