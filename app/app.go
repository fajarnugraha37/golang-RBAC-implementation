package app

import (
	"os"

	"github.com/gin-gonic/gin"

	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/exception"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/guard"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/controller/gift"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/controller/root"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/controller/user"
)

type App struct {
	Router *gin.Engine
}

func NewApp(
	jwtAccessToken guard.AuthHelper,
	guardRoleRepo guard.RoleRepository,
	rootController *root.RootController,
	userController *user.UserController,
	giftController *gift.GiftController,
) *App {
	var (
		router       = gin.Default()
		nonAuthGuard = guard.NonAuthGuard()
		authGuard    = guard.AuthGuard(jwtAccessToken)
		roleGuard    = guard.RoleGuard(guardRoleRepo)
	)

	router.Use(gin.Logger())
	router.Use(gin.CustomRecovery(exception.CustomRecovery))

	router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	var (
		root  = router.Group("")
		users = router.Group("users")
		gifts = router.Group("gifts")
	)

	{
		root.GET("/", rootController.Ping)
		root.POST("/login", nonAuthGuard, rootController.Login)
		root.POST("/register", nonAuthGuard, rootController.Register)
	}

	users.Use(authGuard, roleGuard)
	{
		users.POST("/change-username", userController.ChangeUsername)
		users.POST("/change-password", userController.ChangePassword)
		users.GET("/", userController.GetItems)
		users.POST("/", userController.AddItem)
		users.GET("/:id", userController.GetItemById)
		users.PUT("/:id", userController.UpdateItem)
		users.PATCH("/:id", userController.UpdateItemAttr)
		users.DELETE("/:id", userController.DeleteItem)
	}

	gifts.Use(authGuard, roleGuard)
	{
		gifts.GET("/", giftController.GetItems)
		gifts.POST("/", giftController.AddItem)
		gifts.GET("/:id", giftController.GetItemById)
		gifts.PUT("/:id", giftController.UpdateItem)
		gifts.PATCH("/:id", giftController.UpdateItemAttr)
		gifts.DELETE("/:id", giftController.DeleteItem)
		gifts.POST("/:id/redeem", giftController.ReedemItem)
		gifts.POST("/:id/rating", giftController.RateItem)
	}

	return &App{Router: router}
}

func (app *App) Run() {
	app.Router.Run(":" + os.Getenv("APP_PORT"))
}
