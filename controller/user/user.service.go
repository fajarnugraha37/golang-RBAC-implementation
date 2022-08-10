package user

import "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"

type UserService interface {
	GetPagging(page int, rowsPerPage int, order string, sort string, relation *entity.UserRelation) (*[]entity.User, int64, int)
	GetUserById(id int, relation *entity.UserRelation) *entity.User
	ChangeUsername(id int, newUsername string) *entity.User
	ChangePassword(id int, oldPassword, password, cofirmPassword string) *entity.User
	CrateUser(username string, password string, point int, roles []string) *entity.User
	UpdateUser(id int, username *string, point *int, roles *[]string) *entity.User
	DeleteUser(id int)
}
