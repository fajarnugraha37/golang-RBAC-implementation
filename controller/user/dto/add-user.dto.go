package dto

type AddUserDto struct {
	Username string   `json:"username" binding:"required,min=4,max=16"`
	Password string   `json:"password" binding:"required,min=4,max=24"`
	Point    int      `json:"point" binding:"required,min=1"`
	Roles    []string `json:"roles" binding:"required,dive,oneof=roots admins clients"`
}
