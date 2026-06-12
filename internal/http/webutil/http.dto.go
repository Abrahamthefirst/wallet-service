package custom_http

type ErrorResponse struct {
	Message    string `json:"message"`
	Error      any    `json:"error"`
	StatusCode uint   `json:"statusCode"`
}
