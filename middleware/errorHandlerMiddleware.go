package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorType string

const (
    UNABLE_TO_SAVE         ErrorType = "UNABLE_TO_SAVE"
    UNABLE_TO_FIND_RESOURCE ErrorType = "UNABLE_TO_FIND_RESOURCE"
    UNABLE_TO_READ         ErrorType = "UNABLE_TO_READ"
    UNAUTHORIZED           ErrorType = "UNAUTHORIZED"
    UNABLE_TO_CONNECT_DB    ErrorType ="UNABLE_TO_CONNECT_DB"
)

type CustomError struct {
    Type    ErrorType
    Message string
}

var ErrorMap = map[ErrorType]int{
    UNABLE_TO_SAVE:         http.StatusInternalServerError,
    UNABLE_TO_FIND_RESOURCE: http.StatusNotFound,
    UNABLE_TO_READ:         http.StatusInternalServerError,
    UNAUTHORIZED:           http.StatusUnauthorized,
    UNABLE_TO_CONNECT_DB:   http.StatusInternalServerError,
}

func ErrorHandlerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next() 
        if err, exists := c.Get("error"); exists {
            customError, ok := err.(CustomError)
            if !ok {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
                return
            }
            statusCode, exists := ErrorMap[customError.Type]
            if !exists {
                statusCode = http.StatusInternalServerError
            }
            c.JSON(statusCode, gin.H{"error": customError.Message})
        }
    }
}
