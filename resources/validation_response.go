package resources

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message    string            `json:"message"`
	Validation []ValidationError `json:"validation,omitempty"`
}
