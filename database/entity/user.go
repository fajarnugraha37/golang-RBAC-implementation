package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRelation struct {
	Role   bool
	Redeem bool
	Rating bool
}

type User struct {
	gorm.Model
	Username string  `gorm:"uniqueIndex;type:varchar(255);not null;"`
	Password string  `gorm:"type:text;not null;"`
	Point    int     `gorm:"not null;"`
	Roles    []*Role `gorm:"many2many:user_role;"`
	Redeems  []*Redeem
	Ratings  []*Rating
}

func NewUser(username, password string, point int) (user *User) {
	user = &User{Username: username, Password: password, Point: point}
	return
}

func (user *User) HashPassword() *User {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 4)
	if err != nil {
		panic(err.Error())
	}
	user.Password = string(password)
	return user
}

func (user *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (user *User) GetRoles() (roles []string) {
	if user.Roles == nil {
		return
	}
	for _, v := range user.Roles {
		roles = append(roles, v.Name)
	}
	return
}

func (user *User) AddRole(role string) *User {
	user.Roles = append(user.Roles, &Role{Name: role})
	return user
}
