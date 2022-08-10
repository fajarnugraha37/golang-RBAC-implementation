//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"

	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/guard"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/jwt"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database"

	rgift "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/repository/gift"
	rrating "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/repository/rating"
	rredeem "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/repository/redeem"
	rrole "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/repository/role"
	ruser "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/repository/user"

	sauth "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/service/auth"
	sgift "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/service/gift"
	suser "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/service/user"

	cgift "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/controller/gift"
	croot "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/controller/root"
	cuser "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/controller/user"
)

var WireSet = wire.NewSet(
	jwt.NewJwtNelper,
	rrole.NewRoleRepository,
	rgift.NewGiftRepository,
	rrating.NewRatingRepository,
	rredeem.NewRedeemRepository,
	ruser.NewUserRepository,

	sauth.NewAuthService,
	wire.Bind(new(sauth.AuthRepository), new(*ruser.UserImplRepository)),
	wire.Bind(new(sauth.JwtHelper), new(*jwt.JwtHelper)),

	sgift.NewGiftService,
	wire.Bind(new(sgift.GiftRepository), new(*rgift.GiftImplRepository)),
	wire.Bind(new(sgift.RatingRepository), new(*rrating.RatingImplRepository)),
	wire.Bind(new(sgift.RedeemRepository), new(*rredeem.RedeemImplRepository)),
	wire.Bind(new(sgift.UserRepository), new(*ruser.UserImplRepository)),

	suser.NewUserService,
	wire.Bind(new(suser.UserRepository), new(*ruser.UserImplRepository)),

	cgift.NewGiftController,
	wire.Bind(new(cgift.GiftService), new(*sgift.GiftImplService)),

	croot.NewRootController,
	wire.Bind(new(croot.AuthService), new(*sauth.AuthImplService)),

	cuser.NewUserController,
	wire.Bind(new(cuser.UserService), new(*suser.UserImplService)),

	// middleware
	wire.Bind(new(guard.AuthHelper), new(*jwt.JwtHelper)),
	wire.Bind(new(guard.RoleRepository), new(*rrole.RoleImplRepository)),
)

func InitializeApp() (app *App) {
	wire.Build(
		database.NewDB,
		WireSet,
		NewApp,
	)

	return
}
