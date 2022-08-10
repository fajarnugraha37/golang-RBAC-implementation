package entity

const (
	ROOT_ROLE   = "roots"
	ADMIN_ROLE  = "admins"
	CLIENT_ROLE = "clients"
)

type Role struct {
	Name        string        `gorm:"type:varchar(255);primarykey"`
	Users       []*User       `gorm:"many2many:user_role;"`
	Permissions []*Permission `gorm:"many2many:role_permission;"`
}
