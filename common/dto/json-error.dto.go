package dto

type jsonError struct {
	StatusCode int       `json:"statusCode"`
	Messages   *[]string `json:"messages"`
}

func JErrors(statusCode int, messages ...string) *jsonError {
	return &jsonError{
		StatusCode: statusCode,
		Messages:   &messages,
	}
}
