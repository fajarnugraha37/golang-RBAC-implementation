package dto

type jsonOks[T any] struct {
	StatusCode int       `json:"statusCode"`
	Result     []T       `json:"result"`
	Messages   *[]string `json:"messages"`
}

func (j *jsonOks[any]) SetMessages(messages ...string) *jsonOks[any] {
	j.Messages = &messages
	return j
}

func JOks[T any](statusCode int, result []T) *jsonOks[T] {
	return &jsonOks[T]{
		StatusCode: statusCode,
		Result:     result,
	}
}
