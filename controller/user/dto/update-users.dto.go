package dto

type UpdateUserDto struct {
	Username string   `json:"username" binding:"required,min=4,max=16"`
	Point    int      `json:"point" binding:"required,min=1"`
	Roles    []string `json:"roles" binding:"required,dive,oneof=roots admins clients"`
}
