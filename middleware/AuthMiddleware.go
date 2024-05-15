package middleware

import (
	"G-PROJECT/utils"

	"github.com/gin-gonic/gin"
)

const (
    tokenCookieName = "token"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenCookie, err := c.Cookie(tokenCookieName)
        if err != nil {
            c.Set("error", CustomError{
                Type:    UNAUTHORIZED,
                Message: "Unauthorized: No token provided",
            })
            return
        }

        _, err = utils.VerifyToken(tokenCookie)
        if err != nil {
            c.Set("error", CustomError{
                Type:    UNAUTHORIZED,
                Message: "Unauthorized: Invalid token",
            })
            return
        }

        c.Next()
    }
}
