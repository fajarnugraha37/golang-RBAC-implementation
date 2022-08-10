package jwt

import (
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/exception"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID int      `json:"UserID"`
	Roles  []string `json:"Roles"`
}

func GetUserClaim(c *gin.Context) *Claims {
	user, exists := c.Get("user")
	if !exists {
		panic(exception.NewForbiddenException("Access Forbidden"))
	}
	claims, ok := user.(*Claims)
	if !ok {
		panic(exception.NewForbiddenException("Access Forbidden"))
	}

	return claims
}

func SetUserClaim(c *gin.Context, claims *Claims) {
	c.Set("user", claims)
}
