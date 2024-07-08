package types

type DefaultErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
