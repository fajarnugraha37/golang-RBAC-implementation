package dto

import "mime/multipart"

type UpdateGiftAttrDto struct {
	Name        *string `form:"name" binding:"omitempty"`
	Point       *int    `form:"point" binding:"omitempty"`
	Excerpt     *string `form:"excerpt" binding:"omitempty"`
	Description *string `form:"description" binding:"omitempty"`
	Stock       *int    `form:"stock" binding:"omitempty,gte=0"`

	ImagePath *string
	Image     *multipart.FileHeader `form:"image" binding:"omitempty"`
}
