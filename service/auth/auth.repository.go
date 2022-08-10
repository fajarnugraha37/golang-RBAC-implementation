package auth

import "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"

type AuthRepository interface {
	GetByUsername(username string, relation *entity.UserRelation) (user *entity.User, err error)
	Create(user *entity.User) error
}
