package dto

type AddRatingDto struct {
	Score int `json:"score"  binding:"required,gte=1,lte=5"`
}
