package http

type (
	Error struct {
		Status *Status `json:"status,omitempty"`
	}

	Status struct {
		Message    string `json:"message,omitempty"`
		StatusCode int    `json:"status_code,omitempty"`
	}
)
