package handlers

import (
	"G-PROJECT/database"
	"G-PROJECT/middleware"
	"G-PROJECT/models"
	"G-PROJECT/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
    var err error
    dbInstance, err = database.NewDatabase()
    if err != nil {
        panic(err)
    }
}

func RefreshTokenHandler(c *gin.Context) {
    var user models.User

    if err := c.BindJSON(&user); err != nil {
        c.Set("error", middleware.CustomError{
            Type:    middleware.UNABLE_TO_READ,
            Message: "Unable to read request",
        })
        return
    }

    storedUser,err := dbInstance.GetUserByUsername(user.Username)
    if err == nil {
        c.Set("error", middleware.CustomError{
            Type:    middleware.UNABLE_TO_FIND_RESOURCE,
            Message: "User not found",
        })
        return
    }

    tokenString, err := utils.GenerateToken(*storedUser)
    if err != nil {
        c.Set("error", middleware.CustomError{
            Type:    middleware.UNABLE_TO_SAVE,
            Message: "Unable to generate token",
        })
        return
    }

    c.SetCookie("token", tokenString, int(time.Now().Add(3*time.Minute).Unix()), "/", "", true, true)
    c.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully!"})
}
