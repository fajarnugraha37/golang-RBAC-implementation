package dto

type HeaderAuthorizationDto struct {
	Authorization int `header:"Authorization" binding:"required"`
}
