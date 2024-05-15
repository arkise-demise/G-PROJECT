package handlers

import (
	"G-PROJECT/database"
	"G-PROJECT/middleware"
	"G-PROJECT/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)
func init() {
	var err error
	dbInstance, err = database.NewDatabase()
	if err != nil {
		panic(err) 
	}
}
const (
    imageUploadPath = "./images/"
    maxUploadSize   = 32 << 20
)

func UploadImageHandler(c *gin.Context) {
    tokenCookie, err := c.Request.Cookie("token")
    if err != nil {
        c.Set("error", middleware.CustomError{
            Type:    middleware.UNAUTHORIZED,
            Message: "Unauthorized",
        })
        return
    }

    _, err = utils.VerifyToken(tokenCookie.Value)
    if err != nil {
        c.Set("error", middleware.CustomError{
            Type:    middleware.UNAUTHORIZED,
            Message: "Unauthorized",
        })
        return
    }

    c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)

    err = c.Request.ParseMultipartForm(maxUploadSize)
    if err != nil {
        c.Set("error", middleware.CustomError{
            Type:    middleware.UNABLE_TO_READ,
            Message: err.Error(),
        })
        return
    }

    file, header, err := c.Request.FormFile("images")
    if err != nil {
        c.Set("error", middleware.CustomError{
            Type:    middleware.UNABLE_TO_READ,
            Message: "No image provided",
        })
        return
    }
    defer file.Close()

    var fileExt string

    switch header.Header.Get("Content-Type") {
    case "image/jpeg":
        fileExt = ".jpg"
    case "image/png":
        fileExt = ".png"
    case "image/gif":
        fileExt = ".gif"
    case "image/bmp":
        fileExt = ".bmp"
    default:
        c.Set("error", middleware.CustomError{
            Type:    middleware.UNABLE_TO_READ,
            Message: "Unsupported image type",
        })
        return
    }

    err = os.MkdirAll(imageUploadPath, os.ModePerm)
    if err != nil {
        c.Set("error", middleware.CustomError{
            Type:    middleware.UNABLE_TO_SAVE,
            Message: err.Error(),
        })
        return
    }

    filename := filepath.Join(imageUploadPath, utils.GenerateRequestID()+fileExt)

    newFile, err := os.Create(filename)
    if err != nil {
        c.Set("error", middleware.CustomError{
            Type:    middleware.UNABLE_TO_SAVE,
            Message: err.Error(),
        })
        return
    }
    defer newFile.Close()

    _, err = io.Copy(newFile, file)
    if err != nil {
        c.Set("error", middleware.CustomError{
            Type:    middleware.UNABLE_TO_SAVE,
            Message: err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{"url": filename})
}
func GetImageHandler(c *gin.Context) {
    filename := c.Param("filename")

    imagePath := filepath.Join(imageUploadPath, filename)

    _, err := os.Stat(imagePath)
    if os.IsNotExist(err) {
        c.Set("error",middleware.CustomError{
            Type: middleware.UNABLE_TO_FIND_RESOURCE,
            Message: "Image not found",
        })
        return
    }

    c.File(imagePath)
}