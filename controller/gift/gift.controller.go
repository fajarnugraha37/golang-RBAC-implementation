package gift

import (
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/dto"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/jwt"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/util"
	gdto "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/controller/gift/dto"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"
	"github.com/gin-gonic/gin"
)

type GiftController struct {
	Service GiftService
}

func NewGiftController(GiftService GiftService) *GiftController {
	return &GiftController{Service: GiftService}
}

func (controller *GiftController) GetItems(c *gin.Context) {
	query := util.QueryValidator[gdto.QueryGetGiftsDto](c)
	query.Page = util.Ternary(query.Page == 0, 1, query.Page)
	query.RowsPerPage = util.Ternary(query.RowsPerPage == 0, 20, query.RowsPerPage)
	query.Order = util.Ternary(query.Order == "", "created_at", query.Order)
	query.Sort = util.Ternary(query.Sort == "", "DESC", query.Sort)
	gifts, totalRows, totalPages := controller.Service.GetPagging(query.Page, query.RowsPerPage, query.Order, query.Sort)

	giftsJson := util.Map(*gifts, func(item entity.Gift) *gdto.GiftDto {
		return gdto.NewGiftDto(
			item.ID, item.Name, item.Point, item.Excerpt,
			item.Description, item.Image, item.Stock,
			item.CountRating, item.AvgRating,
			item.CreatedAt, item.UpdatedAt, nil, nil,
		)
	})

	response := dto.
		JPagging(200, giftsJson).
		SetPagging(query.Page, query.RowsPerPage, totalPages, totalRows)

	c.JSON(response.StatusCode, response)
}

func (controller *GiftController) GetItemById(c *gin.Context) {
	var (
		uri  = util.UriValidator[dto.PathIdDto](c)
		gift = controller.Service.GetGiftById(uri.ID)
	)

	response := dto.JOk(200, gdto.NewGiftDto(
		gift.ID, gift.Name, gift.Point, gift.Excerpt,
		gift.Description, gift.Image, gift.Stock,
		gift.CountRating, gift.AvgRating,
		gift.CreatedAt, gift.UpdatedAt, nil, nil,
	))

	c.JSON(response.StatusCode, response)
}

func (controller *GiftController) ReedemItem(c *gin.Context) {
	var (
		claims   = jwt.GetUserClaim(c)
		uri      = util.UriValidator[dto.PathIdDto](c)
		bodyJson = util.JsonValidator[gdto.AddRedeemDto](c)
		redeem   = controller.Service.RedeemGift(claims.UserID, uri.ID, bodyJson.Amount)
	)

	response := dto.JOk(201, gdto.NewRedeemDto(
		redeem.UserID, redeem.GiftID,
		redeem.Amount, redeem.TotalPoint),
	).SetMessages("Gift berhasil di redeem")

	c.JSON(response.StatusCode, response)
}

func (controller *GiftController) RateItem(c *gin.Context) {
	var (
		claims   = jwt.GetUserClaim(c)
		uri      = util.UriValidator[dto.PathIdDto](c)
		bodyJson = util.JsonValidator[gdto.AddRatingDto](c)
		rating   = controller.Service.RateGift(claims.UserID, uri.ID, bodyJson.Score)
	)

	response := dto.JOk(201, gdto.NewRatingDto(rating.UserID, rating.GiftID, rating.Score))

	c.JSON(response.StatusCode, response)
}

func (controller *GiftController) AddItem(c *gin.Context) {
	form := util.BindValidator[gdto.AddGiftDto](c)
	form.ImagePath = util.PublicImageUpload(c, form.Image)
	gift := controller.Service.CreateGift(form.Name, form.Excerpt, form.Description, form.ImagePath, form.Point, form.Stock)

	response := dto.JOk(201, gdto.NewGiftDto(
		gift.ID, gift.Name, gift.Point, gift.Excerpt,
		gift.Description, gift.Image, gift.Stock,
		gift.CountRating, gift.AvgRating,
		gift.CreatedAt, gift.UpdatedAt, nil, nil,
	)).SetMessages("Gift berhasil ditambah")

	c.JSON(response.StatusCode, response)
}

func (controller *GiftController) UpdateItem(c *gin.Context) {
	uri := util.UriValidator[dto.PathIdDto](c)
	form := util.BindValidator[gdto.UpdateGiftDto](c)
	form.ImagePath = util.PublicImageUpload(c, form.Image)
	gift := controller.Service.UpdateGift(uri.ID, &form.Name, &form.Excerpt, &form.Description, &form.ImagePath, &form.Point, &form.Stock)

	response := dto.JOk(201, gdto.NewGiftDto(
		gift.ID, gift.Name, gift.Point, gift.Excerpt,
		gift.Description, gift.Image, gift.Stock,
		gift.CountRating, gift.AvgRating,
		gift.CreatedAt, gift.UpdatedAt, nil, nil,
	)).SetMessages("Gift berhasil diperbaharui")

	c.JSON(response.StatusCode, response)
}

func (controller *GiftController) UpdateItemAttr(c *gin.Context) {
	uri := util.UriValidator[dto.PathIdDto](c)
	form := util.BindValidator[gdto.UpdateGiftAttrDto](c)
	if form.Image != nil {
		imagePath := util.PublicImageUpload(c, form.Image)
		form.ImagePath = &imagePath
	}
	gift := controller.Service.UpdateGift(uri.ID, form.Name, form.Excerpt, form.Description, form.ImagePath, form.Point, form.Stock)

	response := dto.JOk(201, gdto.NewGiftDto(
		gift.ID, gift.Name, gift.Point, gift.Excerpt,
		gift.Description, gift.Image, gift.Stock,
		gift.CountRating, gift.AvgRating,
		gift.CreatedAt, gift.UpdatedAt, nil, nil,
	)).SetMessages("Gift berhasil diperbaharui")

	c.JSON(response.StatusCode, response)
}

func (controller *GiftController) DeleteItem(c *gin.Context) {
	uri := util.UriValidator[dto.PathIdDto](c)
	controller.Service.DeleteGift(uri.ID)

	response := dto.JOk(201, gdto.NewGiftDeletedDto(uri.ID)).
		SetMessages("Data Gift Berhasil Dihapus")

	c.JSON(response.StatusCode, response)
}
