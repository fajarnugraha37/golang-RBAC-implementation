package util

import (
	"strings"

	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/exception"
	"github.com/gin-gonic/gin"
)

func errTransform(err error) []string {
	errors := strings.ReplaceAll(err.Error(), "Key: ", "")
	errors = strings.ReplaceAll(errors, "\n", "")
	return RemoveEmptyStrings(strings.Split(errors, " tag"))
}

func BindValidator[T any](c *gin.Context) (json T) {
	if err := c.ShouldBind(&json); err != nil {
		panic(exception.NewBadRequestException(errTransform(err)...))
	}

	return
}

func JsonValidator[T any](c *gin.Context) (json T) {
	if err := c.ShouldBindJSON(&json); err != nil {
		panic(exception.NewBadRequestException(errTransform(err)...))
	}

	return
}

func UriValidator[T any](c *gin.Context) (json T) {
	if err := c.ShouldBindUri(&json); err != nil {
		panic(exception.NewBadRequestException(errTransform(err)...))
	}

	return
}

func QueryValidator[T any](c *gin.Context) (json T) {
	if err := c.ShouldBindQuery(&json); err != nil {
		panic(exception.NewBadRequestException(errTransform(err)...))
	}

	return
}

func HeaderValidator[T any](c *gin.Context) (json T) {
	if err := c.ShouldBindHeader(&json); err != nil {
		panic(exception.NewBadRequestException(errTransform(err)...))
	}

	return
}
