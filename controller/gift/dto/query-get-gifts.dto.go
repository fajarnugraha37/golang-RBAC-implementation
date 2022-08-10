package dto

type QueryGetGiftsDto struct {
	Page        int    `form:"page"  binding:"omitempty,gte=1"`
	RowsPerPage int    `form:"rows_per_page"  binding:"omitempty,gte=1,lte=25"`
	Order       string `form:"order"  binding:"omitempty,oneof=id name point stock created_at updated_at avg_rating count_rating"`
	Sort        string `form:"sort"  binding:"omitempty,oneof=ASC DESC"`
}
