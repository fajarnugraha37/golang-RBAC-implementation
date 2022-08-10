package gift

import (
	"math"

	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/exception"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/util"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"
)

type GiftImplService struct {
	UserRepo   UserRepository
	GiftRepo   GiftRepository
	RatingRepo RatingRepository
	RedeemRepo RedeemRepository
}

func NewGiftService(userRepo UserRepository, giftRepo GiftRepository, ratingRepo RatingRepository, redeemRepo RedeemRepository) *GiftImplService {
	return &GiftImplService{UserRepo: userRepo, GiftRepo: giftRepo, RatingRepo: ratingRepo, RedeemRepo: redeemRepo}
}

func (service *GiftImplService) RedeemGift(userId int, giftId int, amount int) *entity.Redeem {
	// pastikan user exists
	user, err := service.UserRepo.GetById(userId, nil)
	if err != nil {
		panic(exception.NewBadRequestException("user tidak ditemukan"))
	}

	// pastikan gift exists
	gift, err := service.GiftRepo.GetById(giftId, nil)
	if err != nil {
		panic(exception.NewBadRequestException("gift tidak ditemukan"))
	}

	// pastikan amount > 0 dan kurang dari gift.stock
	if amount > int(gift.Stock) {
		panic(exception.NewBadRequestException("sotck barang tidak memenuhi amount"))
	}

	// pastikan total points <= user points
	totalPoin := amount * int(gift.Point)
	if user.Point < totalPoin {
		panic(exception.NewBadRequestException("point anda tidak memnuhi untuk melakukan redeem"))
	}

	// insert redeem
	redeem := &entity.Redeem{
		UserID:     user.ID,
		GiftID:     gift.ID,
		Amount:     uint(amount),
		TotalPoint: uint(totalPoin),
	}
	err = service.RedeemRepo.Create(redeem)
	util.Panic(err)

	// update point users
	user.Point -= totalPoin
	service.UserRepo.Update(user)

	return redeem
}

func (service *GiftImplService) RateGift(userId int, giftId int, score int) *entity.Rating {
	// validasi score > 0 dan <= 5
	if score <= 0 || score > 5 {
		panic(exception.NewBadRequestException("score harus diatas 0 atau kurang dari sama dengan 5"))
	}

	// pastikan userId exists
	user, err := service.UserRepo.GetById(userId, nil)
	if err != nil {
		panic(exception.NewBadRequestException("user tidak ditemukan"))
	}

	// pastikan giftId exists
	gift, err := service.GiftRepo.GetById(giftId, nil)
	if err != nil {
		panic(exception.NewBadRequestException("gift tidak ditemukan"))
	}

	// upsert data ke db
	rating := &entity.Rating{
		UserID: user.ID,
		GiftID: gift.ID,
		Score:  uint(score),
	}
	err = service.RatingRepo.Upsert(rating)
	util.Panic(err)

	return rating
}

func (service *GiftImplService) GetPagging(page int, rowsPerPage int, order string, sort string) (gifts *[]entity.Gift, totalRows int64, totalPages int) {
	var (
		err    error
		offset = (page - 1) * rowsPerPage
	)

	// get gift items
	gifts, err = service.GiftRepo.GetMany(rowsPerPage, offset, order, sort, &entity.GiftRelation{AvgRating: true})
	util.Panic(err)

	// get gift total count items
	totalRows, err = service.GiftRepo.GetCount()
	util.Panic(err)

	// perhitungan pagging
	totalPages = int(math.Ceil(float64(totalRows) / float64(rowsPerPage)))

	return
}

func (service *GiftImplService) GetGiftById(id int) *entity.Gift {
	// get gift by id
	gift, err := service.GiftRepo.GetById(id, &entity.GiftRelation{AvgRating: true})
	if err != nil {
		panic(exception.NewBadRequestException("data gift tidak ada"))
	}

	return gift
}

func (service *GiftImplService) CreateGift(name, excerpt, description, image string, point, stock int) *entity.Gift {
	// create gift entity
	gift := &entity.Gift{Name: name, Excerpt: excerpt, Description: description, Image: image, Point: uint(point), Stock: uint16(stock)}

	// insert data ke db
	err := service.GiftRepo.Create(gift)
	util.Panic(err)

	return gift
}

func (service *GiftImplService) UpdateGift(id int, name, excerpt, description, image *string, point, stock *int) *entity.Gift {
	// cari gift ada atau tidak
	gift, err := service.GiftRepo.GetById(id, nil)
	if err != nil {
		panic(exception.NewBadRequestException("data user tidak ada"))
	}

	// apply perubahan data
	if name != nil {
		gift.Name = *name
	}
	if excerpt != nil {
		gift.Excerpt = *excerpt
	}
	if description != nil {
		gift.Description = *description
	}
	if image != nil {
		gift.Image = *image
	}
	if point != nil {
		gift.Point = uint(*point)
	}
	if stock != nil {
		gift.Stock = uint16(*stock)
	}

	// update data gift
	err = service.GiftRepo.Update(gift)
	util.Panic(err)

	return gift
}

func (service *GiftImplService) DeleteGift(id int) {
	// delete gift by id
	_, err := service.GiftRepo.Delete(id)
	util.Panic(err)
}
