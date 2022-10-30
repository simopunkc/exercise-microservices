package domain

type MiddlewareError struct {
	Hash  string `json:"hash"`
	Error string `json:"error,omitempty"`
}
