package dto

type RegisterDto struct {
	Username        string `json:"username" binding:"required,min=4,max=16"`
	Password        string `json:"password" binding:"required,eqfield=ConfirmPassword,min=4,max=24"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password,min=4,max=24"`
}
