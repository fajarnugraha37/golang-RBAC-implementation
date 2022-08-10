package dto

type LoginDto struct {
	Username string `json:"username" binding:"required,min=4,max=16"`
	Password string `json:"password" binding:"required,min=4,max=24"`
}
