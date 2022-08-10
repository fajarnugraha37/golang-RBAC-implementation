package guard

import (
	"strings"

	"github.com/gin-gonic/gin"

	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/exception"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/jwt"
)

type AuthHelper interface {
	ValidateToken(token string) (*jwt.Claims, error)
}

func AuthGuard(helper AuthHelper) func(c *gin.Context) {
	return func(c *gin.Context) {
		authorization := strings.Split(c.Request.Header.Get("authorization"), " ")
		if len(authorization) < 2 || strings.ToLower(authorization[0]) != "bearer" {
			panic(exception.NewUnauthorizedException("Unathorized"))
		}
		claims, err := helper.ValidateToken(authorization[1])
		if err != nil {
			panic(exception.NewUnauthorizedException("token expired", err.Error()))
		}
		jwt.SetUserClaim(c, claims)

		c.Next()
	}
}
