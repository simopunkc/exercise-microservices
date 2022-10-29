package domain

type MiddlewareError struct {
	Hash  string `json:"hash"`
	Error error  `json:"error,omitempty"`
}
