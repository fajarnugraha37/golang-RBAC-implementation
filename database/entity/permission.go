package entity

const (
	GET_ACCESS    = "GET"
	POST_ACCESS   = "POST"
	PUT_ACCESS    = "PUT"
	PATCH_ACCESS  = "PATCH"
	DELETE_ACCESS = "DELETE"
)

type Permission struct {
	Name   string  `gorm:"type:varchar(255);primarykey"`
	Path   string  `gorm:"index;type:varchar(255);not null;"`
	Access string  `gorm:"index;type:varchar(255);not null;"`
	Roles  []*Role `gorm:"many2many:role_permission;"`
}
