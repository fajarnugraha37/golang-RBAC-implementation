package dto

import "time"

type UserDto struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Point     int       `json:"point"`
	Roles     []string  `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserDto(id int, username string, point int, roles []string, createdAt, updatedAt time.Time) *UserDto {
	return &UserDto{
		ID: id, Username: username, Point: point, CreatedAt: createdAt, UpdatedAt: updatedAt, Roles: roles,
	}
}
