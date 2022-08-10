package gift

import (
	"strings"

	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/util"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"
	"gorm.io/gorm"
)

type GiftImplRepository struct {
	DB *gorm.DB
}

var repository *GiftImplRepository

func NewGiftRepository(db *gorm.DB) *GiftImplRepository {
	return util.Singleton(repository, func() *GiftImplRepository {
		return &GiftImplRepository{DB: db}
	})
}

func (repo *GiftImplRepository) GetMany(limit, offset int, order, sort string, relation *entity.GiftRelation) (*[]entity.Gift, error) {
	gifts := &[]entity.Gift{}
	query := repo.DB.
		Model(gifts).
		Offset(offset).
		Limit(limit).
		Order(order + " " + strings.ToUpper(sort))
	if relation != nil {
		if relation.Ratings {
			query.Preload("Ratings")
		}
		if relation.Redeems {
			query.Preload("Redeems")
		}
		if relation.AvgRating {
			// jika mod 0.5 itu biarkan else AVG bertindak
			query.Select("gifts.*, avg.avg_rating, avg.count_rating").
				Joins(`
					LEFT JOIN (SELECT gift_id, COUNT(gift_id) as count_rating, IF(AVG(score) % 0.5 = 0, AVG(score), ROUND(AVG(score))) as avg_rating FROM ratings GROUP BY gift_id) as avg 
					ON avg.gift_id = gifts.id
				`)
		}
	}
	err := query.Find(gifts).Error

	return gifts, err
}

func (repo *GiftImplRepository) GetCount() (count int64, err error) {
	err = repo.DB.Table("gifts").Count(&count).Error

	return
}

func (repo *GiftImplRepository) GetById(id int, relation *entity.GiftRelation) (*entity.Gift, error) {
	gift := &entity.Gift{}
	query := repo.DB.
		Model(gift)
	if relation != nil {
		if relation.Ratings {
			query.Preload("Ratings")
		}
		if relation.Redeems {
			query.Preload("Redeems")
		}
		if relation.AvgRating {
			// jika mod 0.5 itu biarkan else AVG bertindak
			query.Select("gifts.*, avg.avg_rating, avg.count_rating").
				Joins(`
				LEFT JOIN (SELECT gift_id, COUNT(gift_id) as count_rating, IF(AVG(score) % 0.5 = 0, AVG(score), ROUND(AVG(score))) as avg_rating FROM ratings GROUP BY gift_id) as avg 
				ON avg.gift_id = gifts.id
			`)
		}
	}
	err := query.First(gift, id).Error

	return gift, err
}

func (repo *GiftImplRepository) Create(gift *entity.Gift) error {
	return repo.DB.Model(gift).Create(gift).Error
}

func (repo *GiftImplRepository) Update(gift *entity.Gift) error {
	return repo.DB.Model(gift).Save(gift).Error
}

func (repo *GiftImplRepository) Delete(id int) (*entity.Gift, error) {
	gift := &entity.Gift{}
	err := repo.DB.Model(gift).Delete(gift, id).Error

	return gift, err
}
