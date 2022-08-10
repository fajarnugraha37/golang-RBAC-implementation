package seeder

import (
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"
)

func GetRoles() *[]entity.Role {
	return &[]entity.Role{
		{
			Name: entity.ROOT_ROLE,
		},
		{
			Name: entity.ADMIN_ROLE,
		},
		{
			Name: entity.CLIENT_ROLE,
		},
	}
}
