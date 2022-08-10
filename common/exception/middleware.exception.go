package exception

import (
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/dto"
	"github.com/gin-gonic/gin"
)

func CustomRecovery(c *gin.Context, recovered interface{}) {
	messages, statusCode := []string{"Internal Server Error"}, 500

	if pointer, ok := instanceOf[CustomException](recovered); ok {
		statusCode = pointer.GetCode()
		messages = pointer.Errors()
	} else if err, ok := recovered.(error); ok {
		messages = []string{err.Error()}
	} else if err, ok := recovered.(string); ok {
		messages = []string{err}
	}

	c.JSON(statusCode, dto.JErrors(statusCode, messages...))
	c.Abort()
}

func instanceOf[T any](objectPtr any) (result *T, ok bool) {
	refrence, okRefrence := objectPtr.(T)
	pointer, okPointer := objectPtr.(*T)
	ok = okPointer || okRefrence
	result = pointer
	if pointer == nil {
		result = &refrence
	}

	return
}
