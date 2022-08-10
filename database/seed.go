package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/seeder"
)

type Seed struct {
	Run func(*gorm.DB) error
}

func SeedAll() []Seed {
	seeds := make([]Seed, 0)

	// roles seed
	seeds = append(seeds, Seed{
		Run: func(d *gorm.DB) error {
			d.Unscoped().Delete(&entity.Role{}, "true")
			return d.
				Clauses(clause.OnConflict{DoNothing: true}).
				Create(seeder.GetRoles()).
				Error
		},
	})

	// permissions seed
	seeds = append(seeds, Seed{
		Run: func(d *gorm.DB) error {
			d.Unscoped().Delete(&entity.Permission{}, "true")
			return d.
				Clauses(clause.OnConflict{DoNothing: true}).
				Create(seeder.GetPermissions()).
				Error
		},
	})

	// user seeds
	seeds = append(seeds, Seed{
		Run: func(d *gorm.DB) error {
			d.Unscoped().Delete(&entity.User{}, "true")
			return d.
				Clauses(clause.OnConflict{DoNothing: true}).
				Create(seeder.GetUsers()).
				Error
		},
	})

	// gifts seeds
	seeds = append(seeds, Seed{
		Run: func(d *gorm.DB) error {
			d.Unscoped().Delete(&entity.Gift{}, "true")
			return d.
				Clauses(clause.OnConflict{DoNothing: true}).
				Create(seeder.GetGifts()).
				Error
		},
	})

	// rating seeds
	seeds = append(seeds, Seed{
		Run: func(d *gorm.DB) error {
			d.Unscoped().Delete(&entity.Rating{}, "true")
			return d.
				Clauses(clause.OnConflict{DoNothing: true}).
				Create(seeder.GetRatings()).Error
		},
	})

	return seeds
}
