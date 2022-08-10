package gift

import "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"

type GiftService interface {
	GetPagging(page int, rowsPerPage int, order string, sort string) (*[]entity.Gift, int64, int)
	GetGiftById(id int) *entity.Gift
	RedeemGift(userId int, giftId int, amount int) *entity.Redeem
	RateGift(userId int, giftId int, score int) *entity.Rating
	CreateGift(name, excerpt, description, image string, point, stock int) *entity.Gift
	UpdateGift(id int, name, excerpt, description, image *string, point, stock *int) *entity.Gift
	DeleteGift(id int)
}
