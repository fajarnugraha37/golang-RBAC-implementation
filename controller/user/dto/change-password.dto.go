package dto

type ChangePasswordDto struct {
	OldPassword        string `json:"old_password" binding:"required,min=4,max=24"`
	NewPassword        string `json:"new_password" binding:"required,eqfield=ConfirmPassword,min=4,max=24"`
	ConfrimNewPassword string `json:"confirm_new_password"  binding:"required,min=4,max=24"`
}
