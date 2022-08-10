package user

import (
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/dto"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/jwt"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/util"
	udto "bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/controller/user/dto"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service UserService
}

func NewUserController(UserService UserService) *UserController {
	return &UserController{Service: UserService}
}

func (controller *UserController) ChangeUsername(c *gin.Context) {
	var (
		claims   = jwt.GetUserClaim(c)
		bodyJson = util.JsonValidator[udto.ChangeUsernameDto](c)
		user     = controller.Service.ChangeUsername(claims.UserID, bodyJson.Username)
	)
	response := dto.JOk(201, udto.NewUserDto(
		int(user.ID), user.Username,
		user.Point, user.GetRoles(),
		user.CreatedAt, user.UpdatedAt,
	)).SetMessages("username berhasil dirubah")

	c.JSON(response.StatusCode, response)
}

func (controller *UserController) ChangePassword(c *gin.Context) {
	var (
		claims   = jwt.GetUserClaim(c)
		bodyJson = util.JsonValidator[udto.ChangePasswordDto](c)
		user     = controller.Service.ChangePassword(claims.UserID, bodyJson.OldPassword, bodyJson.NewPassword, bodyJson.ConfrimNewPassword)
	)

	response := dto.JOk(201, udto.NewUserDto(
		int(user.ID), user.Username,
		user.Point, user.GetRoles(),
		user.CreatedAt, user.UpdatedAt,
	)).SetMessages("password berhasil dirubah")

	c.JSON(response.StatusCode, response)
}

func (controller *UserController) GetItems(c *gin.Context) {
	query := util.QueryValidator[udto.QueryGetUsersDto](c)
	query.Page = util.Ternary(query.Page == 0, 1, query.Page)
	query.RowsPerPage = util.Ternary(query.RowsPerPage == 0, 20, query.RowsPerPage)
	query.Order = util.Ternary(query.Order == "", "created_at", query.Order)
	query.Sort = util.Ternary(query.Sort == "", "DESC", query.Sort)
	users, totalRows, totalPages := controller.Service.GetPagging(query.Page, query.RowsPerPage, query.Order, query.Sort, &entity.UserRelation{Role: true})

	usersJson := util.Map(*users, func(item entity.User) *udto.UserDto {
		return udto.NewUserDto(int(item.ID), item.Username, item.Point, item.GetRoles(), item.CreatedAt, item.UpdatedAt)
	})
	response := dto.JPagging(200, usersJson).
		SetPagging(query.Page, query.RowsPerPage, totalPages, totalRows)

	c.JSON(response.StatusCode, response)
}

func (controller *UserController) GetItemById(c *gin.Context) {
	var (
		uri  = util.UriValidator[dto.PathIdDto](c)
		user = controller.Service.GetUserById(uri.ID, &entity.UserRelation{Role: true})
	)

	response := dto.JOk(200, udto.NewUserDto(
		int(user.ID), user.Username,
		user.Point, user.GetRoles(),
		user.CreatedAt, user.UpdatedAt,
	))

	c.JSON(response.StatusCode, response)
}

func (controller *UserController) AddItem(c *gin.Context) {
	var (
		bodyJson = util.JsonValidator[udto.AddUserDto](c)
		user     = controller.Service.CrateUser(bodyJson.Username, bodyJson.Password, bodyJson.Point, bodyJson.Roles)
	)

	response := dto.JOk(201, udto.NewUserDto(
		int(user.ID), user.Username,
		user.Point, user.GetRoles(),
		user.CreatedAt, user.UpdatedAt,
	)).SetMessages("User berhasil dibuat")

	c.JSON(response.StatusCode, response)
}
func (controller *UserController) UpdateItem(c *gin.Context) {
	var (
		uri      = util.UriValidator[dto.PathIdDto](c)
		bodyJson = util.JsonValidator[udto.UpdateUserDto](c)
		user     = controller.Service.UpdateUser(uri.ID, &bodyJson.Username, &bodyJson.Point, &bodyJson.Roles)
	)

	response := dto.JOk(201, udto.NewUserDto(
		int(user.ID), user.Username,
		user.Point, user.GetRoles(),
		user.CreatedAt, user.UpdatedAt,
	)).SetMessages("Data User Berhasil Diperbaharui")

	c.JSON(response.StatusCode, response)
}

func (controller *UserController) UpdateItemAttr(c *gin.Context) {
	var (
		uri      = util.UriValidator[dto.PathIdDto](c)
		bodyJson = util.JsonValidator[udto.UpdateUserAttrDto](c)
		user     = controller.Service.UpdateUser(uri.ID, bodyJson.Username, bodyJson.Point, bodyJson.Roles)
	)

	response := dto.JOk(201, udto.NewUserDto(
		int(user.ID), user.Username,
		user.Point, user.GetRoles(),
		user.CreatedAt, user.UpdatedAt,
	)).SetMessages("Data User Berhasil Diperbaharui")

	c.JSON(response.StatusCode, response)
}

func (controller *UserController) DeleteItem(c *gin.Context) {
	uri := util.UriValidator[dto.PathIdDto](c)
	controller.Service.DeleteUser(uri.ID)

	response := dto.JOk[any](201, udto.NewUserDeletedDto(uri.ID)).
		SetMessages("Data User Berhasil Dihapus")

	c.JSON(response.StatusCode, response)
}
