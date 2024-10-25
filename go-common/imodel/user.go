package imodel

import "time"

// User defines model for user.
type User struct {
	Created   time.Time `json:"created"`
	Email     string    `json:"email"`
	Id        string    `json:"id"`
	Image     *string   `json:"image,omitempty"`
	LastLogin time.Time `json:"last_login"`
	Name      string    `json:"name"`
}
