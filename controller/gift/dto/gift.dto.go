package dto

import "time"

type GiftDto struct {
	ID          uint         `json:"id"`
	Name        string       `json:"name"`
	Point       uint         `json:"point"`
	Excerpt     string       `json:"excerpt"`
	Description string       `json:"description"`
	Image       string       `json:"image"`
	Stock       uint16       `json:"stock"`
	CountRating *uint        `json:"count_rating"`
	AvgRating   *float32     `json:"avg_rating"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Redeems     []*RedeemDto `json:"redeems"`
	Ratings     []*RatingDto `json:"ratings"`
}

func NewGiftDto(
	id uint, name string, point uint, excerpt, description,
	image string, stock uint16, countRating *uint, avgRating *float32,
	createdAt time.Time, updatedAt time.Time,
	redeems []*RedeemDto, ratings []*RatingDto,
) *GiftDto {
	return &GiftDto{
		ID: id, Name: name, Point: point, Excerpt: excerpt, Description: description,
		Image: image, Stock: stock, CountRating: countRating, AvgRating: avgRating,
		CreatedAt: createdAt, UpdatedAt: updatedAt,
		Redeems: redeems, Ratings: ratings,
	}
}
