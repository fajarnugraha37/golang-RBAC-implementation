package util

import (
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/exception"
	"github.com/gin-gonic/gin"
)

func PublicImageUpload(c *gin.Context, file *multipart.FileHeader) string {
	if file == nil {
		panic(exception.NewBadRequestException("upload file is required"))
	}
	names := strings.Split(file.Filename, ".")
	extensions := names[len(names)-1]
	if !InArray(extensions, []string{"png", "jpg", "jpeg", "svg"}) {
		panic(exception.NewBadRequestException("file upload mst be 'png', 'jpg', 'jpeg' or 'svg'"))
	}

	imageFile := strconv.FormatInt(time.Now().Unix(), 10) + "." + extensions
	imagePath := "./assets/images/" + imageFile
	err := c.SaveUploadedFile(file, imagePath)
	Panic(err)

	return "/assets/images/" + imageFile
}
