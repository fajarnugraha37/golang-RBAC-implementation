package dto

type RatingDto struct {
	UserID uint `json:"user_id"`
	GiftID uint `json:"gift_id"`
	Score  uint `json:"score"`
}

func NewRatingDto(userId, giftId, score uint) *RatingDto {
	return &RatingDto{UserID: userId, GiftID: giftId, Score: score}
}
