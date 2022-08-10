package entity

import "gorm.io/gorm"

type GiftRelation struct {
	Redeems   bool
	Ratings   bool
	AvgRating bool
}

type Gift struct {
	gorm.Model
	Name        string   `gorm:"type:varchar(255);not null;"`
	Point       uint     `gorm:"type:int;not null;"`
	Excerpt     string   `gorm:"type:varchar(255);not null;"`
	Description string   `gorm:"type:text;not null;"`
	Image       string   `gorm:"type:text;not null;"`
	Stock       uint16   `gorm:"type:int;not null;"`
	CountRating *uint    `gorm:"-:migration;<-:false;column:count_rating;"`
	AvgRating   *float32 `gorm:"-:migration;<-:false:column:avg_rating;"`
	Redeems     []*Redeem
	Ratings     []*Rating
}
