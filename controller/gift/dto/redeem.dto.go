package dto

type RedeemDto struct {
	UserID     uint `json:"user_id"`
	GiftID     uint `json:"gift_id"`
	Amount     uint `json:"amount"`
	TotalPoint uint `json:"total_point"`
}

func NewRedeemDto(userId, giftId, amount, totalPoint uint) *RedeemDto {
	return &RedeemDto{UserID: userId, GiftID: giftId, Amount: amount, TotalPoint: totalPoint}
}
