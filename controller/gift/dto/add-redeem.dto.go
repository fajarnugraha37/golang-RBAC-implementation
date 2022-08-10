package dto

type AddRedeemDto struct {
	Amount int `json:"amount"  binding:"required,gte=1"`
}
