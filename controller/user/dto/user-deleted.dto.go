package dto

type UserDeletedDto struct {
	ID int `json:"id"`
}

func NewUserDeletedDto(id int) *UserDeletedDto {
	return &UserDeletedDto{ID: id}
}
