package entity

import (
	"gorm.io/gorm"
)

type RedeemRelation struct {
	User bool
	Gift bool
}

type Redeem struct {
	gorm.Model
	Amount     uint `gorm:"type:int;not null;"`
	TotalPoint uint `gorm:"type:int;not null;"`
	UserID     uint
	User       *User `gorm:"foreignkey:UserID"`
	GiftID     uint
	Gift       *Gift `gorm:"foreignkey:GiftID"`
}
