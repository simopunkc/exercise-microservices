package domain

import "time"

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	NoHp      string    `json:"no_hp"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginParam struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterParam struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,gte=6"`
}

type Repository struct {
	Hash        string `json:"hash,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	User        User   `json:"user,omitempty"`
	RawResponse string `json:"raw_response,omitempty"`
	Error       error  `json:"error,omitempty"`
}

type Service struct {
	Hash        string `json:"hash"`
	User        User   `json:"user,omitempty"`
	RawResponse string `json:"raw_response,omitempty"`
	Error       error  `json:"error,omitempty"`
}

type Handler struct {
	Hash        string `json:"hash"`
	StatusCode  int    `json:"status_code,omitempty"`
	User        *User  `json:"user,omitempty"`
	RawResponse string `json:"raw_response,omitempty"`
	Error       string `json:"error,omitempty"`
}
