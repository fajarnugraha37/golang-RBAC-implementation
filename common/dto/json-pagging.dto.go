package dto

type jsonPagging[T any] struct {
	StatusCode int          `json:"statusCode"`
	Result     []T          `json:"result"`
	Messages   *[]string    `json:"messages"`
	Pagging    *paggingMeta `json:"pagging"`
}

type paggingMeta struct {
	Page        int   `json:"page"`
	PagePerRows int   `json:"pagePerRows"`
	TotalPages  int   `json:"totalPages"`
	TotalRows   int64 `json:"totalRows"`
}

func (j *jsonPagging[any]) SetPagging(page, pagePerRows, totalPages int, totalRows int64) *jsonPagging[any] {
	j.Pagging = &paggingMeta{
		Page:        page,
		PagePerRows: pagePerRows,
		TotalPages:  totalPages,
		TotalRows:   totalRows,
	}
	return j
}

func (j *jsonPagging[any]) SetMessage(messages ...string) *jsonPagging[any] {
	j.Messages = &messages
	return j
}

func JPagging[T any](statusCode int, result []T) *jsonPagging[T] {
	return &jsonPagging[T]{
		StatusCode: statusCode,
		Result:     result,
	}
}
