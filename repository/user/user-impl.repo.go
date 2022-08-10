package user

import (
	"strings"

	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/util"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"
	"gorm.io/gorm"
)

type UserImplRepository struct {
	DB *gorm.DB
}

var repository *UserImplRepository

func NewUserRepository(db *gorm.DB) *UserImplRepository {
	return util.Singleton(repository, func() *UserImplRepository {
		return &UserImplRepository{DB: db}
	})
}

func (repo *UserImplRepository) GetMany(limit, offset int, order, sort string, relation *entity.UserRelation) (*[]entity.User, error) {
	users := &[]entity.User{}
	query := repo.DB.
		Model(users).
		Offset(offset).
		Limit(limit).
		Order(order + " " + strings.ToUpper(sort))
	if relation != nil {
		if relation.Rating {
			query.Preload("Ratings")
		}
		if relation.Redeem {
			query.Preload("Redeems")
		}
		if relation.Role {
			query.Preload("Roles")
		}
	}
	err := query.Find(users).Error

	return users, err
}

func (repo *UserImplRepository) GetCount() (count int64, err error) {
	err = repo.DB.Table("users").Count(&count).Error

	return
}

func (repo *UserImplRepository) GetById(id int, relation *entity.UserRelation) (*entity.User, error) {
	user := &entity.User{}
	query := repo.DB.
		Model(user)
	if relation != nil {
		if relation.Rating {
			query.Preload("Ratings")
		}
		if relation.Redeem {
			query.Preload("Redeems")
		}
		if relation.Role {
			query.Preload("Roles")
		}
	}
	err := query.First(user, id).Error

	return user, err
}

func (repo *UserImplRepository) GetByUsername(username string, relation *entity.UserRelation) (*entity.User, error) {
	user := &entity.User{}
	query := repo.DB.
		Model(user).
		Where("username = ?", username)
	if relation != nil {
		if relation.Rating {
			query.Preload("Ratings")
		}
		if relation.Redeem {
			query.Preload("Redeems")
		}
		if relation.Role {
			query.Preload("Roles")
		}
	}
	err := query.First(user).Error

	return user, err
}

func (repo *UserImplRepository) Create(user *entity.User) error {
	return repo.DB.
		Model(user).
		Create(user).
		Error
}

func (repo *UserImplRepository) Update(user *entity.User) error {
	repo.DB.Model(user).Association("Roles").Replace(user.Roles)
	return repo.DB.Model(user).Save(user).Error
}

func (repo *UserImplRepository) Delete(id int) (*entity.User, error) {
	user := &entity.User{}
	err := repo.DB.
		Model(user).
		Delete(user, id).Error

	return user, err
}
