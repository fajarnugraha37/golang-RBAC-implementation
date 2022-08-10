package auth

import (
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/exception"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/util"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"
)

type AuthImplService struct {
	JwtHelper JwtHelper
	Repo      AuthRepository
}

func NewAuthService(jwtHelper JwtHelper, authRepo AuthRepository) *AuthImplService {
	return &AuthImplService{JwtHelper: jwtHelper, Repo: authRepo}
}

func (service *AuthImplService) Login(username string, password string) (token string) {
	// cari user ada atau tidak
	user, err := service.Repo.GetByUsername(username, &entity.UserRelation{Role: true})
	if err != nil {
		panic(exception.NewBadRequestException("username atau password salah " + err.Error()))
	}

	// validasi apakah password match
	if !user.ValidatePassword(password) {
		panic(exception.NewBadRequestException("username atau password salah"))
	}

	// generate jwt token
	token, err = service.JwtHelper.GenerateToken(int(user.ID), user.GetRoles()...)
	if err != nil {
		panic(err)
	}

	return
}

func (service *AuthImplService) Register(username string, password string, confirmPassword string) string {
	// pastikan usernam tidak ada yang pakai
	_, err := service.Repo.GetByUsername(username, nil)
	if err == nil {
		panic(exception.NewBadRequestException("username telah digunakan pengguna lain, silahkan gunakan username lainnya"))
	}

	// create user dan insert ke database
	user := entity.
		NewUser(username, password, 0).
		AddRole(entity.CLIENT_ROLE).
		HashPassword()
	err = service.Repo.Create(user)
	util.Panic(err)

	// generate jwt token
	token, err := service.JwtHelper.GenerateToken(int(user.ID), user.GetRoles()...)
	util.Panic(err)

	return token
}
