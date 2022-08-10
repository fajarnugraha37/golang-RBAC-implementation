package dto

import "mime/multipart"

type UpdateGiftDto struct {
	Name        string `form:"name" binding:"required"`
	Point       int    `form:"point" binding:"required"`
	Excerpt     string `form:"excerpt" binding:"required"`
	Description string `form:"description" binding:"required"`
	Stock       int    `form:"stock" binding:"required,gte=0"`

	ImagePath string
	Image     *multipart.FileHeader `form:"image" binding:"required"`
}
