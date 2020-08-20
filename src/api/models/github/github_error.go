package github

// ErrorResponse contains info about github api error response
type ErrorResponse struct {
	StatusCode      int     `json:"status_code"`
	Message         string  `json:"message"`
	DocumentationUr string  `json:"documentation_url"`
	Errors          []Error `json:"errors"`
}

func (r ErrorResponse) Error() string {
	return r.Message
}

// Error contains info
type Error struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
