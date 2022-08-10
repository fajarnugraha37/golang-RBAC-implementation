package root

import (
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/dto"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/util"
	rdto "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/controller/root/dto"
	"github.com/gin-gonic/gin"
)

type RootController struct {
	Service AuthService
}

func NewRootController(AuthService AuthService) *RootController {
	return &RootController{Service: AuthService}
}

func (controller *RootController) Ping(c *gin.Context) {
	c.JSON(200, dto.JOk(200, "pong"))
}

func (controller *RootController) Login(c *gin.Context) {
	var (
		bodyJson    = util.JsonValidator[rdto.LoginDto](c)
		accessToken = controller.Service.Login(bodyJson.Username, bodyJson.Password)
	)

	response := dto.
		JOk(201, rdto.NewAuthDto(accessToken)).
		SetMessages("Login Berhasil")

	c.JSON(201, response)
}

func (controller *RootController) Register(c *gin.Context) {
	var (
		bodyJson    = util.JsonValidator[rdto.RegisterDto](c)
		accessToken = controller.Service.Register(bodyJson.Username, bodyJson.Password, bodyJson.ConfirmPassword)
	)

	response := dto.
		JOk(201, rdto.NewAuthDto(accessToken)).
		SetMessages("Registrasi Berhasil")

	c.JSON(201, response)
}
