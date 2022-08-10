package dto

type ChangeUsernameDto struct {
	Username string `json:"username" binding:"required,min=4,max=16"`
}
