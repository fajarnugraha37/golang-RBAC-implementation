package exception

import (
	"net/http"
	"strings"
)

type CustomException struct {
	code     int
	messages []string
}

func NewBadRequestException(messages ...string) *CustomException {
	return &CustomException{code: http.StatusBadRequest, messages: messages}
}

func NewUnauthorizedException(messages ...string) *CustomException {
	return &CustomException{code: http.StatusUnauthorized, messages: messages}
}

func NewForbiddenException(messages ...string) *CustomException {
	return &CustomException{code: http.StatusForbidden, messages: messages}
}

func NewNotFoundException(messages ...string) *CustomException {
	return &CustomException{code: http.StatusNotFound, messages: messages}
}

func NewNotAcceptableException(messages ...string) *CustomException {
	return &CustomException{code: http.StatusNotAcceptable, messages: messages}
}

func NewPreconditionFailedException(messages ...string) *CustomException {
	return &CustomException{code: http.StatusPreconditionFailed, messages: messages}
}

func NewCustomExceptions(code int, messages ...string) *CustomException {
	return &CustomException{code: code, messages: messages}
}

func (m *CustomException) Error() string {
	return strings.Join(m.messages, "\n")
}

func (m *CustomException) Errors() []string {
	return m.messages
}

func (m *CustomException) GetCode() int {
	return m.code
}
