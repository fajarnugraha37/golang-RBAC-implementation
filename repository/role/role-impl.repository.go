package role

import (
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/util"
	"gorm.io/gorm"
)

type RoleImplRepository struct {
	DB *gorm.DB
}

var repository *RoleImplRepository

func NewRoleRepository(db *gorm.DB) *RoleImplRepository {
	return util.Singleton(repository, func() *RoleImplRepository {
		return &RoleImplRepository{DB: db}
	})
}

func (repo *RoleImplRepository) GetRoles(method, path string) []string {
	result := []map[string]interface{}{}
	err := repo.DB.Table("permissions").
		Select("role_permission.role_name as role").
		Where("permissions.access = ? AND permissions.path = ?", method, path).
		Joins("LEFT JOIN role_permission ON role_permission.permission_name = permissions.name").Scan(&result).Error

	roles := []string{}
	if err == nil {
		for _, v := range result {
			item := v["role"].(string)
			roles = append(roles, item)
		}
	}

	return roles
}
