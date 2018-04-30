package models

// Health data to verify the server is working.
type Health struct {
	Active bool `json:"active"`
}

type Status struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
