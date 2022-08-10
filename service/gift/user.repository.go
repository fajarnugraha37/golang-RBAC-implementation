package gift

import "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"

type UserRepository interface {
	GetById(id int, relation *entity.UserRelation) (*entity.User, error)
	Update(user *entity.User) error
}
