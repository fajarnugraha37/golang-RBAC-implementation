package guard

import (
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/exception"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/jwt"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/util"
	"github.com/gin-gonic/gin"
)

type RoleRepository interface {
	GetRoles(method, path string) []string
}

func RoleGuard(repo RoleRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		method, path := c.Request.Method, c.FullPath()
		guranteds := repo.GetRoles(method, path)
		if len(guranteds) < 1 {
			c.Next()
			return
		}
		claims := jwt.GetUserClaim(c)
		if intersect := util.Intersection(claims.Roles, guranteds); len(intersect) < 1 {
			panic(exception.NewForbiddenException("You are not allow to access this url"))
		}

		c.Next()
	}
}
