package dto

type jsonOk[T any] struct {
	StatusCode int       `json:"statusCode"`
	Result     T         `json:"result"`
	Messages   *[]string `json:"messages"`
}

func (j *jsonOk[any]) SetMessages(messages ...string) *jsonOk[any] {
	j.Messages = &messages
	return j
}

func JOk[T any](statusCode int, result T) *jsonOk[T] {
	return &jsonOk[T]{
		StatusCode: statusCode,
		Result:     result,
	}
}
