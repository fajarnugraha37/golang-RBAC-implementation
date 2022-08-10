package dto

type UpdateUserAttrDto struct {
	Username *string   `json:"username" binding:"omitempty,min=4,max=16"`
	Point    *int      `json:"point" binding:"omitempty,min=1"`
	Roles    *[]string `json:"roles" binding:"omitempty,dive,oneof=roots admins clients"`
}
