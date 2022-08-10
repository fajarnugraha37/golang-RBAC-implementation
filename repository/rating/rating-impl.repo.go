package rating

import (
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/util"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RatingImplRepository struct {
	DB *gorm.DB
}

var repository *RatingImplRepository

func NewRatingRepository(db *gorm.DB) *RatingImplRepository {
	return util.Singleton(repository, func() *RatingImplRepository {
		return &RatingImplRepository{DB: db}
	})
}

func (repo *RatingImplRepository) Upsert(rating *entity.Rating) error {
	err := repo.DB.
		Model(rating).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "user_id"},
				{Name: "gift_id"},
			},
			DoUpdates: clause.AssignmentColumns([]string{"score"}),
		}).
		Create(rating).Error

	return err
}
