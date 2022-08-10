package user

import (
	"math"

	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/exception"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/util"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"
)

type UserImplService struct {
	Repo UserRepository
}

func NewUserService(userRepo UserRepository) *UserImplService {
	return &UserImplService{Repo: userRepo}
}

func (service *UserImplService) ChangeUsername(id int, newUsername string) *entity.User {
	// cari user ada atau tidak
	user, err := service.Repo.GetById(id, nil)
	if err != nil {
		panic(exception.NewBadRequestException("data user tidak ada"))
	}

	// username existing dan yang mau diganti engga sama
	if user.Username == newUsername {
		panic(exception.NewBadRequestException("username lama dan username baru masih sama"))
	}

	// cari username sudah dipakai atau tidak
	_, err = service.Repo.GetByUsername(newUsername, nil)
	if err == nil {
		panic(exception.NewBadRequestException("Username sudah dipakai, silahkan gunakan username yang lain"))
	}

	// rubah username
	user.Username = newUsername
	err = service.Repo.Update(user)
	util.Panic(err)

	return user
}

func (service *UserImplService) ChangePassword(id int, oldPassword, password, cofirmPassword string) *entity.User {
	// password lama dan baru sama
	if oldPassword == password {
		panic(exception.NewBadRequestException("Password baru dan konfirmasi password baru tidak sama"))
	}

	// cocokan password dan confirm password
	if password != cofirmPassword {
		panic(exception.NewBadRequestException("Password baru dan konfirmasi password baru tidak sama"))
	}

	// caru user ada atau tidak
	user, err := service.Repo.GetById(id, nil)
	if err != nil {
		panic(exception.NewBadRequestException("data user tidak ada"))
	}

	// apakah oldpassword cocok dengan existing password
	isValid := user.ValidatePassword(oldPassword)
	if !isValid {
		panic(exception.NewBadRequestException("Password lama salah"))
	}

	user.Password = password
	user.HashPassword()
	err = service.Repo.Update(user)
	util.Panic(err)

	return user
}

func (service *UserImplService) CrateUser(username string, password string, point int, roles []string) *entity.User {
	// pastikan username belum dipakai
	_, err := service.Repo.GetByUsername(username, nil)
	if err == nil {
		panic(exception.NewBadRequestException("Username sudah dipakai, silahkan gunakan username yang lain"))
	}

	// create user
	var user = entity.NewUser(username, password, point).HashPassword()
	for _, role := range roles {
		user.AddRole(role)
	}

	// insert user ke db
	service.Repo.Create(user)

	return user
}

func (service *UserImplService) UpdateUser(id int, username *string, point *int, roles *[]string) *entity.User {
	// cari user ada atau tidak
	user, err := service.Repo.GetById(id, nil)
	if err != nil {
		panic(exception.NewBadRequestException("data user tidak ada"))
	}

	// apply perubahan data
	if username != nil {
		// check username exist atau tidak
		_, err = service.Repo.GetByUsername(*username, nil)
		if err == nil {
			panic(exception.NewBadRequestException("Username sudah dipakai, silahkan gunakan username yang lain"))
		}
		user.Username = *username
	}
	if point != nil {
		user.Point = *point
	}
	if roles != nil {
		for _, role := range *roles {
			user.AddRole(role)
		}
	}

	// update data user
	err = service.Repo.Update(user)
	util.Panic(err)

	return user
}

func (service *UserImplService) DeleteUser(id int) {
	// delete user by id
	_, err := service.Repo.Delete(id)
	util.Panic(err)
}

func (service *UserImplService) GetPagging(page int, rowsPerPage int, order string, sort string, relation *entity.UserRelation) (users *[]entity.User, totalRows int64, totalPages int) {
	var (
		err    error
		offset = (page - 1) * rowsPerPage
	)

	// get user items
	users, err = service.Repo.GetMany(rowsPerPage, offset, order, sort, relation)
	util.Panic(err)

	// get user total count items
	totalRows, err = service.Repo.GetCount()
	util.Panic(err)

	// perhitungan pagging
	totalPages = int(math.Ceil(float64(totalRows) / float64(rowsPerPage)))

	return
}

func (service *UserImplService) GetUserById(id int, relation *entity.UserRelation) *entity.User {
	// get user by id
	user, err := service.Repo.GetById(id, relation)
	if err != nil {
		panic(exception.NewBadRequestException("data user tidak ada"))
	}

	return user
}
