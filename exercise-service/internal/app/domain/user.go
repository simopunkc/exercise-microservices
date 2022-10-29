package domain

import "time"

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	NoHP      string    `json:"no_hp"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RepositoryUser struct {
	Hash        string `json:"hash,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	User        User   `json:"user,omitempty"`
	RawResponse string `json:"raw_response,omitempty"`
	Error       error  `json:"error,omitempty"`
}
