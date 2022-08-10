package redeem

import (
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/util"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"
	"gorm.io/gorm"
)

type RedeemImplRepository struct {
	DB *gorm.DB
}

var repository *RedeemImplRepository

func NewRedeemRepository(db *gorm.DB) *RedeemImplRepository {
	return util.Singleton(repository, func() *RedeemImplRepository {
		return &RedeemImplRepository{DB: db}
	})
}

func (repo *RedeemImplRepository) Create(redeem *entity.Redeem) error {
	return repo.DB.Model(redeem).Create(redeem).Error
}
