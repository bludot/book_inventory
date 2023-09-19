package http_utils

type HTTPError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e *HTTPError) Error() string {
	return e.Message
}
