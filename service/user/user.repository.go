package user

import "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"

type UserRepository interface {
	GetById(id int, relation *entity.UserRelation) (*entity.User, error)
	GetByUsername(username string, relation *entity.UserRelation) (*entity.User, error)
	Update(user *entity.User) error
	GetMany(limit, offset int, order, sort string, relation *entity.UserRelation) (*[]entity.User, error)
	GetCount() (int64, error)
	Delete(id int) (*entity.User, error)
	Create(user *entity.User) error
}
