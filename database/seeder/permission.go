package seeder

import (
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"
)

var permissions []entity.Permission = []entity.Permission{
	{
		Name:   "get_users",
		Path:   "/users/",
		Access: "GET",
		Roles: []*entity.Role{
			{Name: entity.ROOT_ROLE},
			{Name: entity.ADMIN_ROLE},
		},
	},
	{
		Name:   "craete_user",
		Path:   "/users/",
		Access: "POST",
		Roles: []*entity.Role{
			{Name: entity.ROOT_ROLE},
		},
	},
	{
		Name:   "get_user_by_id",
		Path:   "/users/:id",
		Access: "GET",
		Roles: []*entity.Role{
			{Name: entity.ROOT_ROLE},
		},
	},
	{
		Name:   "put_user_by_id",
		Path:   "/users/:id",
		Access: "PUT",
		Roles: []*entity.Role{
			{Name: entity.ROOT_ROLE},
		},
	},
	{
		Name:   "patch_user_by_id",
		Path:   "/users/:id",
		Access: "PATCH",
		Roles: []*entity.Role{
			{Name: entity.ROOT_ROLE},
		},
	},
	{
		Name:   "delete_user_by_id",
		Path:   "/users/:id",
		Access: "DELETE",
		Roles: []*entity.Role{
			{Name: entity.ROOT_ROLE},
		},
	},
	{
		Name:   "create_gift",
		Path:   "/gifts/",
		Access: "POST",
		Roles: []*entity.Role{
			{Name: entity.ROOT_ROLE},
		},
	},
	{
		Name:   "put_gift",
		Path:   "/gifts/:id",
		Access: "PUT",
		Roles: []*entity.Role{
			{Name: entity.ROOT_ROLE},
			{Name: entity.ADMIN_ROLE},
		},
	},
	{
		Name:   "patch_gift",
		Path:   "/gifts/:id",
		Access: "PATCH",
		Roles: []*entity.Role{
			{Name: entity.ROOT_ROLE},
			{Name: entity.ADMIN_ROLE},
		},
	},
	{
		Name:   "delete_gift",
		Path:   "/gifts/:id",
		Access: "DELETE",
		Roles: []*entity.Role{
			{Name: entity.ROOT_ROLE},
			{Name: entity.ADMIN_ROLE},
		},
	},
}

func GetPermissions() []entity.Permission {
	return permissions
}
