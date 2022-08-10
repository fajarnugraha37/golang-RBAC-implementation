package entity

import (
	"time"

	"gorm.io/gorm"
)

type RatingRelation struct {
	User bool
	Gift bool
}

type Rating struct {
	// gorm.Model
	UserID    uint  `gorm:"primarykey"`
	User      *User `gorm:"foreignkey:UserID"`
	GiftID    uint  `gorm:"primarykey"`
	Gift      *Gift `gorm:"foreignkey:GiftID"`
	Score     uint  `gorm:"type:int;not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
