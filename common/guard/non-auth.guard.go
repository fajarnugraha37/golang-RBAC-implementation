package guard

import (
	"github.com/gin-gonic/gin"

	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/exception"
)

func NonAuthGuard() func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Header.Get("authorization") != "" {
			panic(exception.NewBadRequestException("you are already logged in"))
		}
		c.Next()
	}
}
