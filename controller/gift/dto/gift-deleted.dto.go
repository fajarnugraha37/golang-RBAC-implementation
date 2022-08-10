package dto

type GiftDeletedDto struct {
	ID int `json:"id"`
}

func NewGiftDeletedDto(id int) *GiftDeletedDto {
	return &GiftDeletedDto{ID: id}
}
